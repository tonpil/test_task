package presenter

import (
	"test_task/app/api"
	"test_task/interactor/dto"
)

type PostNewsPresenter struct {
}

func (p *PostNewsPresenter) Present(outputDTO *dto.PostNewsOutputDTO) *api.PostNewsResponse {
	if outputDTO == nil {
		return nil
	}

	response := &api.PostNewsResponse{
		Success: true,
	}
	return response
}
