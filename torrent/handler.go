package torrent

import (
	"context"
	"errors"
	"fmt"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/types/infohash"
)

type Handler struct {
	torrentClient *torrent.Client
	torrentFiles  map[metainfo.Hash]*torrent.Torrent
}

func NewHandler() *Handler {
	return &Handler{
		torrentFiles: make(map[metainfo.Hash]*torrent.Torrent),
	}
}

func (h *Handler) Init(ctx context.Context) error {
	cfg := torrent.NewDefaultClientConfig()
	// TODO: change this config (maybe to temp folder)
	cfg.DataDir = "/Users/lap14897/Desktop/torrent"
	cfg.Debug = false

	client, err := torrent.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("create torrent client: %w", err)
	}
	h.torrentClient = client

	return nil
}

func (h *Handler) Close(ctx context.Context) error {
	if h.torrentClient != nil {
		errs := h.torrentClient.Close()
		return errors.Join(errs...)
	}
	return nil
}

func infoHashFromHexString(infoHashHex string) (metainfo.Hash, error) {
	var infoHash infohash.T
	err := infoHash.FromHexString(infoHashHex)
	return infoHash, err
}
