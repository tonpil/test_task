package controller

import (
	"context"
	"test_task/config"
)

type NewsControllers struct {
	ctx      context.Context
	dbConfig config.DBConfig
}

func New(ctx context.Context, dbConfig config.DBConfig) *NewsControllers {
	return &NewsControllers{ctx, dbConfig}
}
