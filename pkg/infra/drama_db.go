package infra

import (
	"context"

	"github.com/sakuradaman/go_simple_api/pkg/domain/model"
	"github.com/sakuradaman/go_simple_api/pkg/domain/repository"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type dramaRepository struct {
	db *gorm.DB
}

func NewDramaRepository(db *gorm.DB) repository.DramaRepository {
	return &dramaRepository{db}
}

func (r *dramaRepository) SelectAllDramas(ctx context.Context) ([]*model.Drama, error) {
	var records []*model.Drama
	if result := r.db.Find(&records); result.Error != nil {
		return nil, xerrors.Errorf("Find err %w", result.Error)
	}
	return records, nil
}
