package data

import (
	"context"
	"github.com/golang-learning/inventory/internal/biz"
	"github.com/golang-learning/inventory/internal/data/ent"
	"github.com/pkg/errors"
)

type storageRepo struct {
	client *ent.Client
}

func NewStorageRepo(client *ent.Client) biz.IStorageRepo {
	return &storageRepo{client: client}
}

func (s *storageRepo) GetStorage(ctx context.Context, id int) (*biz.Storage, error) {
	storage, err := s.client.Storage.Get(ctx, id)
	if ent.IsNotFound(err) {
		return nil, errors.Wrapf(err, "user not found, hostname: %d, err: %+v", id, err)
	}

	if err != nil {
		return nil, errors.Wrapf(err, "db query err: %+v, hostname: %d", err, id)
	}

	return &biz.Storage{
		Hostname: storage.Hostname,
		Size:     storage.Size,
		ID:       storage.ID,
	}, nil

}
