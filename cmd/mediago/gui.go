package main

import (
	"context"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	wailsruntime "github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/Sophomoresty/mediago/internal/util"
)

//go:embed all:frontend/dist
var assets embed.FS

// App exposes download operations to the Wails frontend.
type App struct {
	ctx           context.Context
	downloadMutex sync.Mutex
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) emitLog(msg string) {
	wailsruntime.EventsEmit(a.ctx, "log", msg)
}

func (a *App) emitProgress(msg string) {
	wailsruntime.EventsEmit(a.ctx, "progress", msg)
}

// GetDefaultDownloadDir returns <exe dir>/downloads.
func (a *App) GetDefaultDownloadDir() string {
	exePath, err := os.Executable()
	if err != nil {
		return "downloads"
	}
	return filepath.Join(filepath.Dir(exePath), "downloads")
}

// SelectDownloadDir opens the native directory picker.
// Returns "" when the user cancels.
func (a *App) SelectDownloadDir(current string) (string, error) {
	opts := wailsruntime.OpenDialogOptions{Title: "选择下载目录"}
	if st, err := os.Stat(current); err == nil && st.IsDir() {
		opts.DefaultDirectory = current
	}
	return wailsruntime.OpenDirectoryDialog(a.ctx, opts)
}

// StartDownload downloads one URL. Logs stream to the frontend via the
// "log" event and per-file progress via the "progress" event. Downloads
// are serialized: the frontend disables its button while one is running.
func (a *App) StartDownload(url, format, downloadDir, cookies, proxyURL string, allPlaylist bool) error {
	a.downloadMutex.Lock()
	defer a.downloadMutex.Unlock()

	url = strings.TrimSpace(url)
	if url == "" {
		return fmt.Errorf("请输入视频链接")
	}

	downloadDir = strings.TrimSpace(downloadDir)
	if downloadDir == "" {
		downloadDir = a.GetDefaultDownloadDir()
	}
	if err := os.MkdirAll(downloadDir, 0o755); err != nil {
		return fmt.Errorf("无法创建下载目录: %w", err)
	}

	// Empty string clears any proxy left over from a previous download.
	proxyURL = strings.TrimSpace(proxyURL)
	if err := util.SetDefaultProxy(proxyURL); err != nil {
		return fmt.Errorf("代理地址无效: %w", err)
	}

	// Route CLI logging and progress into the frontend for this download.
	guiProgressCallback = a.emitLog
	currentVideoProgressCallback = a.emitProgress
	defer func() {
		guiProgressCallback = nil
		currentVideoProgressCallback = nil
	}()

	// GUI mode never runs cobra, so the flag globals still hold zero
	// values; set the same defaults the CLI flags declare.
	formatSpec = format
	if strings.TrimSpace(formatSpec) == "" {
		formatSpec = "best"
	}
	cookieFile = strings.TrimSpace(cookies)
	cookieBrowser = ""
	proxy = proxyURL
	outputTemplate = filepath.Join(downloadDir, "%(title)s.%(ext)s")
	noProgress = true
	concurrency = 10
	mergeOutputFmt = "mp4"
	downloadAll = allPlaylist
	listFormats = false
	dumpJSON = false
	simulate = false

	start := time.Now()
	a.emitLog(fmt.Sprintf("[开始] %s\n", start.Format("15:04:05")))

	err := processURL(context.Background(), url)

	a.emitLog(fmt.Sprintf("[完成] %s (耗时: %s)\n", time.Now().Format("15:04:05"), time.Since(start).Round(time.Second)))

	if err != nil {
		a.emitLog(fmt.Sprintf("[失败] %v\n", err))
		a.emitLog("[提示] 请检查:\n  1. 网络连接是否正常\n  2. 视频链接是否有效\n  3. 是否需要提供 Cookies\n\n")
		return err
	}

	a.emitLog("[成功] 下载完成\n")
	a.emitLog(fmt.Sprintf("[位置] %s\n\n", downloadDir))
	return nil
}

// runGUI launches the Wails desktop window; called from main when the
// program starts without CLI arguments.
func runGUI() {
	app := &App{}

	err := wails.Run(&options.App{
		Title:     "MediaGo 视频下载器",
		Width:     860,
		Height:    720,
		MinWidth:  520,
		MinHeight: 420,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 245, G: 246, B: 248, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "启动 GUI 失败:", err)
		os.Exit(1)
	}
}
