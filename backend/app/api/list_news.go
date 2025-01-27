package api

import "test_task/domain"

type ListNewsResponse struct {
	News  []domain.News `json:"News"`
	After *string       `json:"After"`
}
