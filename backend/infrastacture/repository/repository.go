package repository

import (
	"context"
	"database/sql"
	"fmt"
	"test_task/config"
	dbentities "test_task/infrastacture/repository/entities"
	"test_task/interactor/iface"

	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type Repository struct {
	db *reform.DB
}

func NewRepository(cfg *config.DBConfig) (*Repository, error) {
	sqlDB, err := sql.Open("postgres", cfg.DBURL)
	fmt.Println(cfg.DBURL)
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("unable to ping database: %v", err)
	}

	db := reform.NewDB(sqlDB, postgresql.Dialect, nil)

	return &Repository{db: db}, nil
}

func (r *Repository) GetNewsByID(ctx context.Context, id int64) (*dbentities.News, error) {
	item, err := r.db.FindOneFrom(dbentities.NewsTable, "id", id)

	if err != nil {
		return nil, TransformError(err)
	}

	return item.(*dbentities.News), nil
}

func (r *Repository) ListNews(ctx context.Context, args iface.ListNewsArgs) ([]dbentities.News, error) {
	var tail string
	var err error
	var items []reform.Struct

	if args.After != nil {
		tail = fmt.Sprintf("WHERE id > %d ORDER BY id LIMIT %d", args.After.ID, args.Limit)
		items, err = r.db.SelectAllFrom(dbentities.NewsTable, tail)
	} else {
		tail = fmt.Sprintf("ORDER BY id LIMIT %d", args.Limit)
		items, err = r.db.SelectAllFrom(dbentities.NewsTable, tail)
	}

	if err != nil {
		return nil, TransformError(err)
	}

	result := make([]dbentities.News, len(items))
	for i, item := range items {
		result[i] = *(item.(*dbentities.News))
	}

	return result, nil
}

func (r *Repository) ListCategoriesByNewsIDs(ctx context.Context, newsIDs []int) ([]dbentities.NewsCategories, error) {
	items := []reform.Struct{}

	for _, newsID := range newsIDs {
		newsCategories, err := r.db.FindAllFrom(dbentities.NewsCategoriesTable, "news_id", newsID)
		if err != nil {
			return nil, TransformError(err)
		}

		items = append(items, newsCategories...)
	}

	result := make([]dbentities.NewsCategories, len(items))
	for i, item := range items {
		result[i] = *(item.(*dbentities.NewsCategories))
	}

	return result, nil
}

func (r *Repository) ListCategoriesByNewsID(ctx context.Context, newsID int64) ([]dbentities.NewsCategories, error) {
	items, err := r.db.SelectAllFrom(dbentities.NewsCategoriesTable, "news_id = ?", newsID)
	if err != nil {
		return nil, TransformError(err)
	}

	result := make([]dbentities.NewsCategories, len(items))
	for i, item := range items {
		result[i] = *(item.(*dbentities.NewsCategories))
	}

	return result, nil
}

func (r *Repository) UpdateNews(ctx context.Context, news dbentities.News) error {
	err := r.db.Save(&news)
	return TransformError(err)
}

func (r *Repository) DeleteCategoriesByNewsID(ctx context.Context, newsID int64) error {
	_, err := r.db.DeleteFrom(dbentities.NewsCategoriesTable, fmt.Sprintf("WHERE news_id = %d", newsID))

	return TransformError(err)
}

func (r *Repository) InsertCategoriesByNewsID(ctx context.Context, args iface.InsertCategoriesByNewsIDArgs) error {
	var insertStatements []reform.Struct

	// Для каждой категории из args.Category создаем строку для вставки
	for _, categoryID := range args.Categories {
		// Добавляем новый объект типа NewsCategories в срез insertStatements
		insertStatements = append(insertStatements, &dbentities.NewsCategories{
			NewsID:     args.NewsID,
			CategoryID: categoryID,
		})
	}
	err := r.db.InsertMulti(insertStatements...)

	if err != nil {
		return TransformError(err)
	}

	return nil

}
