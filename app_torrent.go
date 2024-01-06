package main

import (
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/chrisngyn/torplayer/torrent"
)

func (a *App) AddTorrentFromString(torrentString string) (string, error) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("AddTorrentFromString: %s", torrentString))
	return a.torrentHandler.AddTorrentFromString(a.ctx, torrentString)
}

func (a *App) AddTorrentFromFileContent(content []byte) (string, error) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("AddTorrentFromFileContent with lenght %d", len(content)))
	return a.torrentHandler.AddTorrentFromFileContent(a.ctx, content)
}

func (a *App) GetTorrentInfo(infoHashHex string) (torrent.Info, error) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("GetFiles: %s", infoHashHex))
	return a.torrentHandler.GetTorrentInfo(a.ctx, infoHashHex)
}

func (a *App) StartDownload(infoHashHex, path string) error {
	runtime.LogDebug(a.ctx, fmt.Sprintf("StartDownload: %s %s", infoHashHex, path))
	return a.torrentHandler.StartDownload(a.ctx, infoHashHex, path)
}

func (a *App) StopDownload(infoHashHex, path string) error {
	runtime.LogDebug(a.ctx, fmt.Sprintf("StopDownload: %s %s", infoHashHex, path))
	return a.torrentHandler.StopDownload(a.ctx, infoHashHex, path)
}
