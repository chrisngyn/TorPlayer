package main

import (
	"context"
	"log"

	"github.com/anacrolix/torrent"
)

// App struct
type App struct {
	ctx context.Context

	torrentClient *torrent.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx

	cfg := torrent.NewDefaultClientConfig()
	// TODO: change this config (maybe to temp folder)
	cfg.DataDir = "/Users/lap14897/Desktop/torrent"
	cfg.Debug = true

	client, err := torrent.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to create torrent client: %s", err)
	}
	a.torrentClient = client

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
	if a.torrentClient != nil {
		errs := a.torrentClient.Close()
		for err := range errs {
			log.Printf("Failed to close torrent client: %v", err)
		}
	}
}
