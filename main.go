package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

//go:embed web/*
var webFiles embed.FS

func main() {
	// Set up HTTP server
	mux := http.NewServeMux()

	// Serve embedded web files
	mux.Handle("/", http.FileServer(http.FS(webFiles)))

	// API endpoint for downloading
	mux.HandleFunc("/api/download", handleDownload)

	// API endpoint for getting supported sites
	mux.HandleFunc("/api/sites", handleGetSites)

	// Start server
	port := "8080"
	addr := fmt.Sprintf("127.0.0.1:%s", port)
	url := fmt.Sprintf("http://%s/web/", addr)

	// Start server in background
	go func() {
		log.Printf("Starting server on %s", addr)
		if err := http.ListenAndServe(addr, mux); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait a moment for server to start
	time.Sleep(500 * time.Millisecond)

	// Open browser
	log.Printf("Opening browser: %s", url)
	if err := openBrowser(url); err != nil {
		log.Printf("Failed to open browser: %v", err)
		log.Printf("Please manually open: %s", url)
	} else {
		log.Printf("Browser opened successfully")
	}

	// Keep server running
	log.Println("Server is running. Press Ctrl+C to stop.")
	select {}
}

func handleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	format := r.FormValue("format")
	cookies := r.FormValue("cookies")
	proxy := r.FormValue("proxy")

	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	result, err := executeMediago(url, format, cookies, proxy)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v\n\n%s", err, result)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, result)
}

func handleGetSites(w http.ResponseWriter, r *http.Request) {
	sites, err := getExtractors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, sites)
}

func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default: // linux, freebsd, openbsd, netbsd
		cmd = exec.Command("xdg-open", url)
	}

	// Don't wait for the command to finish
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Start()
}
