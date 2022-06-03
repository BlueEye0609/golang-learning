package usecase

import (
	"context"
	"github.com/golang-learning/inventory/internal/biz"
)

type storage struct {
	repo biz.IStorageRepo
}

func NewStorageUsecase(repo biz.IStorageRepo) biz.IStorageUseCase {
	return &storage{repo: repo}
}

// GetUserInfo GetUserInfo
func (s *storage) GetStorage(ctx context.Context, id int) (*biz.Storage, error) {
	return s.repo.GetStorage(ctx, id)
}
