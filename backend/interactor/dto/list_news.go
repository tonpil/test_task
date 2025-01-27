package dto

import (
	"test_task/domain"
)

type ListNewsInputDTO struct {
	After *domain.NewsCursor
	Limit *int
}

type ListNewsOutputDTO struct {
	Items []domain.News
	After *domain.NewsCursor
}
