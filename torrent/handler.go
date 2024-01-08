package torrent

import (
	"context"
	"errors"
	"fmt"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/types/infohash"
)

type Config struct {
	DataDir string `json:"dataDir"`
}

func (c Config) Validate() error {
	if c.DataDir == "" {
		return errors.New("data dir is required")
	}
	return nil
}

type Handler struct {
	appCtx        context.Context
	torrentClient *torrent.Client
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Init(ctx context.Context, config Config) error {
	if err := config.Validate(); err != nil {
		return fmt.Errorf("validate config: %w", err)
	}
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = config.DataDir
	cfg.Debug = false

	client, err := torrent.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("create torrent client: %w", err)
	}

	h.appCtx = ctx
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
