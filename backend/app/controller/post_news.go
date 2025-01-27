package controller

import (
	"fmt"
	"strconv"
	"test_task/app/api"
	"test_task/app/presenter"
	"test_task/infrastacture/repository"
	"test_task/interactor/dto"
	usecase "test_task/interactor/use_case"

	"github.com/gofiber/fiber"
)

// @Summary PostNews
// @Tags News
// @Accept json
// @Produce json
// @Param Id path int true "News ID"
// @Param request body api.PostNewsRequest true "Request for edit news"
// @Success 200 {object} api.PostNewsResponse
// @Router /edit/{Id} [post]
func (s *NewsControllers) PostNews(c *fiber.Ctx) {

	idStr := c.Params("Id")
	var request api.PostNewsRequest

	if err := c.BodyParser(&request); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if idStr == "" || err != nil {
		c.Status(fiber.StatusBadRequest).JSON(api.ErrorResponse{
			Message: "Invalid Id format",
			Code:    fiber.StatusBadRequest,
		})
		return
	}

	presenter := presenter.PostNewsPresenter{}
	inputDTO := dto.PostNewsInputDTO{
		ID:          id,
		Title:       request.Title,
		Content:     request.Content,
		CategoryIDs: request.CategoryIDs,
	}

	repository, err := repository.NewRepository(&s.dbConfig)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
			Message: fmt.Sprintf("Internal server error: %v", err),
			Code:    fiber.StatusInternalServerError,
		})
		return
	}

	useCase := usecase.PostNewsUseCase{
		Repository: repository,
	}

	outputDTO, err := useCase.Execute(s.ctx, inputDTO)

	if err != nil {
		code, err := TransformErrorToHttpError(err)
		c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
			Message: fmt.Sprintf("Internal server error: %v", err),
			Code:    code,
		})
		return
	}

	response := presenter.Present(outputDTO)

	c.Status(fiber.StatusOK).JSON(response)
}
