package main

import (
	"context"
	"github.com/golang-learning/inventory/internal/data"
	"github.com/golang-learning/inventory/internal/data/ent"
	"github.com/golang-learning/inventory/internal/service"
	"github.com/golang-learning/inventory/internal/usecase"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

func InitStorageService() (*service.StorageService, error) {
	v, err := InitConfig()
	if err != nil {
		return nil, err
	}

	client, err := NewDB(v)
	if err != nil {
		return nil, err
	}

	iStorageRepo := data.NewStorageRepo(client)
	iStorageUsecase := usecase.NewStorageUsecase(iStorageRepo)
	storageService := service.NewStorageService(iStorageUsecase)
	return storageService, nil
}

var StorageSet = wire.NewSet(
	service.NewStorageService,
	data.NewStorageRepo,
	usecase.NewStorageUsecase,
)

func NewDB(v *viper.Viper) (*ent.Client, error) {
	client, err := ent.Open(
		v.Sub("db").GetString("type"),
		v.Sub("db").GetString("dsn"),
	)

	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}

func InitConfig() (*viper.Viper, error) {
	viper.AddConfigPath("../configs")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return viper.GetViper(), nil
}
