package slog

import (
	"log/slog"

	"github.com/sakuradaman/go_simple_api/pkg/lib/config"
	"golang.org/x/xerrors"
)

// slog の設定
func SetUp(c *config.Config) error {
	// ログレベルの設定
	if err := slog.SetDefaultLevel(slog.LevelDebug); err != nil {
		return xerrors.Errorf("fail to set default level: %w", err)
	}
}
