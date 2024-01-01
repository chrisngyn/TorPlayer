package main

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/types/infohash"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context

	torrentClient *torrent.Client
	torrentFiles  map[metainfo.Hash]*torrent.Torrent
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
	a.torrentFiles = make(map[metainfo.Hash]*torrent.Torrent)

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

func (a *App) AddTorrentFromString(torrentString string) (string, error) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("AddTorrentFromString: %s", torrentString))
	// TODO: Implement
	return "", nil
}

func (a *App) AddTorrentFromFileContent(content []byte) (string, error) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("AddTorrentFromFileContent with lenght %d", len(content)))
	info, err := metainfo.Load(bytes.NewReader(content))
	if err != nil {
		return "", fmt.Errorf("load metainfo: %s", err)
	}

	tor, err := a.torrentClient.AddTorrent(info)
	if err != nil {
		return "", fmt.Errorf("add torrent: %s", err)
	}

	a.torrentFiles[tor.InfoHash()] = tor

	return tor.InfoHash().String(), nil
}

type TorrentInfo struct {
	Name   string `json:"name"`
	Length int64  `json:"length"`
	Files  []File `json:"files"`
}

type File struct {
	DisplayPath    string `json:"displayPath"`
	Length         int64  `json:"length"`
	BytesCompleted int64  `json:"bytesCompleted"`
}

func (a *App) GetTorrentInfo(infoHashHex string) (TorrentInfo, error) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("GetFiles: %s", infoHashHex))
	var infoHash infohash.T
	err := infoHash.FromHexString(infoHashHex)
	if err != nil {
		return TorrentInfo{}, fmt.Errorf("parse infohash: %s", err)
	}
	tor, ok := a.torrentFiles[infoHash]
	if !ok {
		return TorrentInfo{}, fmt.Errorf("torrent not found")
	}
	<-tor.GotInfo()

	torInfo := TorrentInfo{
		Name:   tor.Name(),
		Length: tor.Length(),
		Files:  make([]File, 0),
	}

	for _, file := range tor.Files() {
		torInfo.Files = append(torInfo.Files, File{
			DisplayPath:    file.DisplayPath(),
			Length:         file.Length(),
			BytesCompleted: file.BytesCompleted(),
		})
	}
	return torInfo, nil
}
