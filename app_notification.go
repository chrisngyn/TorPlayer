package main

import (
	"github.com/chrisngyn/torplayer/notification"
)

func (a *App) NotifyAlert(title, message string) {
	notification.Alert(title, message)
}
