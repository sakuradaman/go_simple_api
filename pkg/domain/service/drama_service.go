package service

import (
	"context"

	"github.com/sakuradaman/go_simple_api/pkg/domain/model"
	"github.com/sakuradaman/go_simple_api/pkg/domain/repository"
)

// インターフェース
type DramaService interface {
	SelectAllDramas(ctx context.Context) ([]*model.Drama, error)
}

// インターフェースを実装した構造体
type dramaService struct {
	dramaRepo repository.DramaRepository
}

// repoインターフェースからserviceインターフェースを作成
func NewDramaService(dramaRepo repository.DramaRepository) DramaService {
	return &dramaService{
		dramaRepo: dramaRepo,
	}
}

// インターフェースのメソッドを実装
func (s *dramaService) SelectAllDramas(ctx context.Context) ([]*model.Drama, error) {
	return s.dramaRepo.SelectAllDramas(ctx)
}
