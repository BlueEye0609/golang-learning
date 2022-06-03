package biz

import "context"

type Storage struct {
	Hostname string
	Size     int
	ID       int
}

type IStorageRepo interface {
	GetStorage(ctx context.Context, id int) (*Storage, error)
}

type IStorageUseCase interface {
	GetStorage(ctx context.Context, id int) (*Storage, error)
}
