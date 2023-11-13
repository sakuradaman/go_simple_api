package usecase

import (
	"context"

	"github.com/sakuradaman/go_simple_api/pkg/domain/model"
	"github.com/sakuradaman/go_simple_api/pkg/domain/service"
)

type DramaUsecase interface {
	SelectAllDramas(ctx context.Context) ([]*model.Drama, error)
}

type dramaUsecase struct {
	svc service.DramaService
}

// serviceインターフェースからusecaseインターフェースを作成
func NewDramaUsecase(ds service.DramaService) DramaUsecase {
	return &dramaUsecase{
		svc: ds,
	}
}

func (du *dramaUsecase) SelectAllDramas(ctx context.Context) ([]*model.Drama, error) {
	mDrama, err := du.svc.SelectAllDramas(ctx)
	if err != nil {
		return nil, err
	}
	return mDrama, nil
}
