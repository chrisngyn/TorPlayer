package main

import (
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/chrisngyn/torplayer/subtitle"
)

func (a *App) StandardizeSubtitle(content []byte, fileExtension string, addedMilliseconds int64) (out []byte, err error) {
	runtime.LogDebug(a.ctx, fmt.Sprintf("StandardizeSubtitle: %s", fileExtension))
	return subtitle.Normalize(content, fileExtension, time.Duration(addedMilliseconds)*time.Millisecond)
}
