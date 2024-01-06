package main

import (
	"context"
	"log"

	"github.com/chrisngyn/torplayer/torrent"
)

// App struct
type App struct {
	ctx context.Context

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

	if err := a.torrentHandler.Init(ctx); err != nil {
		log.Fatalf("Failed to initialize torrent handler: %v", err)
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
		log.Fatalf("Failed to close torrent handler: %v", err)
	}
}
