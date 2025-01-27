package presenter

import (
	"encoding/base64"
	"fmt"
	"test_task/app/api"
	"test_task/interactor/dto"
)

type ListNewsPresenter struct {
}

func (p *ListNewsPresenter) Present(outputDTO *dto.ListNewsOutputDTO) *api.ListNewsResponse {
	if outputDTO == nil || outputDTO.Items == nil {
		return nil
	}

	var after *string

	if outputDTO.After != nil {
		afterIDStr := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", outputDTO.After.ID)))
		after = &afterIDStr
	}

	response := &api.ListNewsResponse{
		News:  outputDTO.Items,
		After: after,
	}
	return response
}
