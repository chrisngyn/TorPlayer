package main

import (
	"github.com/chrisngyn/torplayer/torrent"
)

func (a *App) GetCurrentConfig() torrent.Config {
	return a.config
}
