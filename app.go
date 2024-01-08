package main

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/chrisngyn/torplayer/torrent"
)

// App struct
type App struct {
	ctx context.Context

	config         torrent.Config
	torrentHandler *torrent.Handler
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		torrentHandler: torrent.NewHandler(),
	}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	// Get the store directory
	storeDirectory, err := os.UserHomeDir()
	if err != nil {
		runtime.LogFatalf(ctx, "Failed to get user home directory: %v", err)
	}
	storeDirectory = storeDirectory + "/.torplayer"
	// Create the store directory if it doesn't exist
	if _, err := os.Stat(storeDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(storeDirectory, 0700); err != nil {
			runtime.LogFatalf(ctx, "Failed to create store directory: %v", err)
		}
	}

	a.config = torrent.Config{
		DataDir: storeDirectory,
	}

	if err := a.torrentHandler.Init(ctx, a.config); err != nil {
		runtime.LogFatalf(ctx, "Failed to initialize torrent handler: %v", err)
	}
}

// domReady is called after the front-end dom has been loaded
func (a *App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue,
// false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	if err := a.torrentHandler.Close(ctx); err != nil {
		runtime.LogErrorf(ctx, "Failed to close torrent handler: %v", err)
	}

	if err := os.RemoveAll(a.config.DataDir); err != nil {
		runtime.LogErrorf(ctx, "Failed to remove store directory: %v", err)
	}
}
