package iface

import (
	"context"
	"test_task/domain"
	dbentities "test_task/infrastacture/repository/entities"
)

type Repository interface {
	GetNewsByID(ctx context.Context, id int64) (*dbentities.News, error)
	ListNews(ctx context.Context, args ListNewsArgs) ([]dbentities.News, error)
	ListCategoriesByNewsIDs(ctx context.Context, newsIDs []int) ([]dbentities.NewsCategories, error)
	UpdateNews(ctx context.Context, news dbentities.News) error
	InsertCategoriesByNewsID(ctx context.Context, args InsertCategoriesByNewsIDArgs) error
	DeleteCategoriesByNewsID(ctx context.Context, newsID int64) error
}

type ListNewsArgs struct {
	After *domain.NewsCursor
	Limit int32
}

type UpdateNewsArgs struct {
	ID      int64
	Title   string
	Content string
}

type InsertCategoriesByNewsIDArgs struct {
	NewsID     int64
	Categories []int64
}
