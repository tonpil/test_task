package usecase

import (
	"context"
	"test_task/domain"
	"test_task/interactor/dto"
	"test_task/interactor/iface"
)

type ListNewsUseCase struct {
	Repository iface.Repository
}

const limitListNews = 10

func (u *ListNewsUseCase) Execute(ctx context.Context, inputDTO dto.ListNewsInputDTO) (*dto.ListNewsOutputDTO, error) {
	var limitForDB int32

	if inputDTO.Limit != nil {
		limitForDB = int32(*inputDTO.Limit) + 1
	} else {
		limitForDB = limitListNews
	}

	news, err := u.Repository.ListNews(ctx, iface.ListNewsArgs{
		After: inputDTO.After,
		Limit: limitForDB,
	})

	if err != nil {
		return nil, err
	}

	newsIDs := make([]int, len(news))
	for i, item := range news {
		newsIDs[i] = int(item.ID)
	}

	newsCategories, err := u.Repository.ListCategoriesByNewsIDs(ctx, newsIDs)

	if err != nil {
		return nil, err
	}

	newsIDToCategories := make(map[int64][]int64)
	for _, newsCategory := range newsCategories {
		newsIDToCategories[newsCategory.NewsID] = append(newsIDToCategories[newsCategory.NewsID], newsCategory.CategoryID)
	}

	result := make([]domain.News, len(news))
	for i, item := range news {
		result[i] = domain.News{
			ID:       item.ID,
			Title:    item.Title,
			Content:  item.Content,
			Category: newsIDToCategories[item.ID],
		}
	}

	var after *domain.NewsCursor

	if int32(len(result)) == limitForDB {
		after = &domain.NewsCursor{
			ID: result[limitForDB-2].ID,
		}
		result = result[:limitForDB-1]
	} else {
		after = nil
	}

	return &dto.ListNewsOutputDTO{
		After: after,
		Items: result,
	}, nil
}
