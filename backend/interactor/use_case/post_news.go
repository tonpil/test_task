package usecase

import (
	"context"
	"test_task/interactor/dto"
	"test_task/interactor/iface"

	dbentities "test_task/infrastacture/repository/entities"
)

type PostNewsUseCase struct {
	Repository iface.Repository
}

func (u *PostNewsUseCase) Execute(ctx context.Context, inputDTO dto.PostNewsInputDTO) (*dto.PostNewsOutputDTO, error) {
	news, err := u.Repository.GetNewsByID(ctx, inputDTO.ID)

	if err != nil {
		return nil, err
	}

	var titleToDB, textToDB string

	if inputDTO.Title != nil {
		titleToDB = *inputDTO.Title
	} else {
		titleToDB = news.Title
	}

	if inputDTO.Content != nil {
		textToDB = *inputDTO.Content
	} else {
		textToDB = news.Content
	}

	err = u.Repository.UpdateNews(ctx, dbentities.News{
		ID:      news.ID,
		Title:   titleToDB,
		Content: textToDB,
	})

	if err != nil {
		return nil, err
	}

	if inputDTO.CategoryIDs != nil {
		err = u.Repository.DeleteCategoriesByNewsID(ctx, inputDTO.ID)

		if err != nil {
			return nil, err
		}

		err = u.Repository.InsertCategoriesByNewsID(ctx, iface.InsertCategoriesByNewsIDArgs{
			NewsID:     inputDTO.ID,
			Categories: inputDTO.CategoryIDs,
		})

		if err != nil {
			return nil, err
		}
	}

	return &dto.PostNewsOutputDTO{}, nil
}
