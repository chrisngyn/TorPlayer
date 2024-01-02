package torrent

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

func (h *Handler) AddTorrentFromString(ctx context.Context, torrentString string) (string, error) {
	var (
		tor *torrent.Torrent
		err error
	)

	if strings.HasPrefix(torrentString, "magnet:") {
		tor, err = h.torrentClient.AddMagnet(torrentString)
		if err != nil {
			return "", fmt.Errorf("add magnet: %w", err)
		}
	} else if strings.HasPrefix(torrentString, "http") || strings.HasPrefix(torrentString, "https") {
		resp, err := http.Get(torrentString)
		if err != nil {
			return "", fmt.Errorf("get torrent file: %w", err)
		}
		defer resp.Body.Close()
		info, err := metainfo.Load(resp.Body)
		if err != nil {
			return "", fmt.Errorf("load metainfo: %w", err)
		}
		tor, err = h.torrentClient.AddTorrent(info)
		if err != nil {
			return "", fmt.Errorf("add torrent: %w", err)
		}
	} else {
		infoHash, err := infoHashFromHexString(strings.ToLower(torrentString))
		if err != nil {
			return "", fmt.Errorf("invalid info hash: %w", err)
		}
		tor, _ = h.torrentClient.AddTorrentInfoHash(infoHash)
	}

	return tor.InfoHash().String(), nil
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

	return tor.InfoHash().String(), nil
}
