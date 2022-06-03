package service

import (
	"context"
	v1 "github.com/golang-learning/inventory/api/storage/v1"
	"github.com/golang-learning/inventory/internal/biz"
	"google.golang.org/grpc/metadata"
	"strconv"
)

type StorageService struct {
	v1.StorageServiceServer
	usecase biz.IStorageUseCase
}

func NewStorageService(usecase biz.IStorageUseCase) *StorageService {
	return &StorageService{usecase: usecase}
}

func (s *StorageService) GetStorage(ctx context.Context, req *v1.GetRequest) (*v1.StorageReply, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	data := md.Get("uid")

	id, _ := strconv.Atoi(data[0])

	storage, err := s.usecase.GetStorage(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &v1.StorageReply{
		Hostname: storage.Hostname,
		Size:     int64(storage.Size),
	}

	return resp, nil
}