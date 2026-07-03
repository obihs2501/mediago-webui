package main

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "MediaGo WebUI",
		Width:  1200,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 247, G: 244, B: 239, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal("Error:", err)
	}
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// DownloadVideo starts a video download
func (a *App) DownloadVideo(url string, format string, cookies string, proxy string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("URL is required")
	}

	// Call mediago binary
	result, err := executeMediago(url, format, cookies, proxy)
	if err != nil {
		return "", err
	}

	return result, nil
}

// GetSupportedSites returns list of supported sites
func (a *App) GetSupportedSites() (string, error) {
	return getExtractors()
}
