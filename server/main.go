package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	tasks      = make(map[string]*DownloadTask)
	tasksMutex sync.RWMutex
	clients    = make(map[*websocket.Conn]bool)
	clientsMux sync.RWMutex
	broadcast  = make(chan TaskUpdate)
)

type DownloadTask struct {
	ID          string    `json:"id"`
	URL         string    `json:"url"`
	Status      string    `json:"status"` // pending, downloading, completed, failed
	Progress    float64   `json:"progress"`
	Title       string    `json:"title"`
	Output      string    `json:"output"`
	Error       string    `json:"error,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	cmd         *exec.Cmd
}

type TaskUpdate struct {
	Type string        `json:"type"`
	Task *DownloadTask `json:"task"`
}

type DownloadRequest struct {
	URL            string `json:"url"`
	Format         string `json:"format,omitempty"`
	Output         string `json:"output,omitempty"`
	Cookies        string `json:"cookies,omitempty"`
	CookiesBrowser string `json:"cookies_browser,omitempty"`
	Proxy          string `json:"proxy,omitempty"`
	YesPlaylist    bool   `json:"yes_playlist,omitempty"`
}

type ExtractorInfo struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	NeedAuth bool   `json:"need_auth"`
}

func main() {
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// UI
	r.HandleFunc("/", serveIndex).Methods("GET")

	// API endpoints
	r.HandleFunc("/api/download", handleDownload).Methods("POST")
	r.HandleFunc("/api/tasks", handleListTasks).Methods("GET")
	r.HandleFunc("/api/tasks/{id}", handleGetTask).Methods("GET")
	r.HandleFunc("/api/tasks/{id}/cancel", handleCancelTask).Methods("POST")
	r.HandleFunc("/api/extractors", handleListExtractors).Methods("GET")
	r.HandleFunc("/ws", handleWebSocket)

	// Start broadcast handler
	go handleBroadcasts()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting MediaGo WebUI on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/index.html")
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	var req DownloadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	taskID := fmt.Sprintf("task-%d", time.Now().UnixNano())
	task := &DownloadTask{
		ID:        taskID,
		URL:       req.URL,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	tasksMutex.Lock()
	tasks[taskID] = task
	tasksMutex.Unlock()

	// Start download in background
	go executeDownload(task, req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func executeDownload(task *DownloadTask, req DownloadRequest) {
	task.Status = "downloading"
	notifyClients(task)

	// Find mediago binary
	mediagoPath := findMediagoBinary()
	if mediagoPath == "" {
		task.Status = "failed"
		task.Error = "mediago binary not found"
		notifyClients(task)
		return
	}

	// Build command
	args := []string{}

	if req.Format != "" {
		args = append(args, "-f", req.Format)
	}

	if req.Output != "" {
		args = append(args, "-o", req.Output)
	} else {
		args = append(args, "-o", "downloads/%(title)s.%(ext)s")
	}

	if req.Cookies != "" {
		args = append(args, "--cookies", req.Cookies)
	}

	if req.CookiesBrowser != "" {
		args = append(args, "--cookies-from-browser", req.CookiesBrowser)
	}

	if req.Proxy != "" {
		args = append(args, "--proxy", req.Proxy)
	}

	if req.YesPlaylist {
		args = append(args, "--yes-playlist")
	}

	args = append(args, "--write-info-json")
	args = append(args, req.URL)

	// Create downloads directory
	os.MkdirAll("downloads", 0755)

	// Execute command
	cmd := exec.Command(mediagoPath, args...)
	task.cmd = cmd

	output, err := cmd.CombinedOutput()

	task.CompletedAt = time.Now()

	if err != nil {
		task.Status = "failed"
		task.Error = fmt.Sprintf("%s\nOutput: %s", err.Error(), string(output))
	} else {
		task.Status = "completed"
		task.Progress = 100
		task.Output = string(output)

		// Try to extract title from output
		if title := extractTitle(string(output)); title != "" {
			task.Title = title
		}
	}

	notifyClients(task)
}

func findMediagoBinary() string {
	// Check current directory
	paths := []string{
		"./mediago.exe",
		"./mediago",
		"../mediago.exe",
		"../mediago",
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			abs, _ := filepath.Abs(path)
			return abs
		}
	}

	// Check PATH
	if path, err := exec.LookPath("mediago"); err == nil {
		return path
	}

	return ""
}

func extractTitle(output string) string {
	// Simple extraction - look for common patterns
	// This is a placeholder - real implementation would parse output better
	return ""
}

func handleListTasks(w http.ResponseWriter, r *http.Request) {
	tasksMutex.RLock()
	defer tasksMutex.RUnlock()

	taskList := make([]*DownloadTask, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskList)
}

func handleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	tasksMutex.RLock()
	task, exists := tasks[taskID]
	tasksMutex.RUnlock()

	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func handleCancelTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	tasksMutex.Lock()
	task, exists := tasks[taskID]
	tasksMutex.Unlock()

	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	if task.cmd != nil && task.cmd.Process != nil {
		task.cmd.Process.Kill()
		task.Status = "cancelled"
		task.Error = "Cancelled by user"
		notifyClients(task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func handleListExtractors(w http.ResponseWriter, r *http.Request) {
	mediagoPath := findMediagoBinary()
	if mediagoPath == "" {
		http.Error(w, "mediago binary not found", http.StatusInternalServerError)
		return
	}

	cmd := exec.Command(mediagoPath, "--list-extractors")
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	clientsMux.Lock()
	clients[conn] = true
	clientsMux.Unlock()

	defer func() {
		clientsMux.Lock()
		delete(clients, conn)
		clientsMux.Unlock()
	}()

	// Keep connection alive
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}

func notifyClients(task *DownloadTask) {
	update := TaskUpdate{
		Type: "task_update",
		Task: task,
	}
	broadcast <- update
}

func handleBroadcasts() {
	for update := range broadcast {
		clientsMux.RLock()
		for client := range clients {
			err := client.WriteJSON(update)
			if err != nil {
				log.Printf("WebSocket write error: %v", err)
				client.Close()
				clientsMux.Lock()
				delete(clients, client)
				clientsMux.Unlock()
			}
		}
		clientsMux.RUnlock()
	}
}
