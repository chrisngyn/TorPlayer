package torrent

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/anacrolix/torrent"
	"github.com/gabriel-vasile/mimetype"
)

func (h *Handler) StartDownload(ctx context.Context, infoHashHex, path string) error {
	file, err := h.getFile(infoHashHex, path)
	if err != nil {
		return fmt.Errorf("get file: %w", err)
	}

	file.Download()
	return nil
}

func (h *Handler) StopDownload(ctx context.Context, infoHashHex, path string) error {
	file, err := h.getFile(infoHashHex, path)
	if err != nil {
		return fmt.Errorf("get file: %w", err)
	}

	file.SetPriority(torrent.PiecePriorityNone)
	return nil
}

func (h *Handler) getFile(infoHashHex, path string) (*torrent.File, error) {
	infoHash, err := infoHashFromHexString(infoHashHex)
	if err != nil {
		return nil, fmt.Errorf("parse infohash: %w", err)
	}

	tor, ok := h.torrentClient.Torrent(infoHash)
	if !ok {
		return nil, fmt.Errorf("torrent not found")
	}

	var file *torrent.File
	for _, f := range tor.Files() {
		if f.DisplayPath() == path {
			file = f
			break
		}
	}

	if file == nil {
		return nil, fmt.Errorf("file not found")
	}

	return file, nil
}

// ServeHTTP serves the file content, it is used to stream the file content to the client.
// If there is anything error, should return http.StatusNotFound, otherwise the client will
// be stuck in the loading state. (I don't know why)
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// url will be have format /stream/:infoHash/file/:filePath
	// e.g. /stream/1234567890abcdef/file/path/to/file.mp4
	if strings.HasPrefix(r.URL.Path, "/stream/") == false {
		http.Error(w, "invalid url", http.StatusNotFound)
	}
	tokens := strings.SplitN(strings.Trim(r.URL.Path, "/"), "/", 4)
	if len(tokens) != 4 {
		http.Error(w, "invalid url", http.StatusNotFound)
		return
	}
	infoHashHex, filePath := tokens[1], tokens[3]

	log.Printf("infoHashHex: %s", infoHashHex)
	log.Printf("filePath: %s", filePath)

	file, err := h.getFile(infoHashHex, filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// TODO: Maybe implement more effective file reader to seed up the download
	reader := file.NewReader()
	reader.SetResponsive()
	reader.SetReadahead(file.Length() / 100) // Read ahead 1% of the file

	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		log.Printf("Detect mime type fail: %v", err)
	} else {
		log.Printf("Mime type: %s", mime.String())
		w.Header().Set("Content-Type", mime.String())
	}

	http.ServeContent(w, r, file.DisplayPath(), time.Now(), reader)
}
