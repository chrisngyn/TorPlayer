package torrent

import (
	"bytes"
	"context"
	"fmt"

	"github.com/anacrolix/torrent/metainfo"
)

func (h *Handler) AddTorrentFromString(ctx context.Context, torrentString string) (string, error) {
	// TODO: Implement
	return "", nil
}

func (h *Handler) AddTorrentFromFileContent(ctx context.Context, content []byte) (string, error) {
	info, err := metainfo.Load(bytes.NewReader(content))
	if err != nil {
		return "", fmt.Errorf("load metainfo: %w", err)
	}

	tor, err := h.torrentClient.AddTorrent(info)
	if err != nil {
		return "", fmt.Errorf("add torrent: %w", err)
	}

	h.torrentFiles[tor.InfoHash()] = tor

	return tor.InfoHash().String(), nil
}
