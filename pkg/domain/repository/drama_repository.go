package repository

import (
	"context"

	"github.com/sakuradaman/go_simple_api/pkg/domain/model"
)

// インターフェース
type DramaRepository interface {
	SelectAllDramas(ctx context.Context) ([]*model.Drama, error)
}
