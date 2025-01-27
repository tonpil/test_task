package controller

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"test_task/app/api"
	"test_task/app/presenter"
	"test_task/domain"
	"test_task/infrastacture/repository"
	"test_task/interactor/dto"
	usecase "test_task/interactor/use_case"

	"github.com/gofiber/fiber"
)

// @Summary ListNews
// @Description Retrieve a list of news
// @Tags News
// @Accept json
// @Produce json
// @Param limit query int false "Count of news"
// @Param after query string false "Pagination string"
// @Success 200 {object} api.ListNewsResponse
// @Router /list [get]
func (s *NewsControllers) ListNews(c *fiber.Ctx) {

	limitStr := c.Query("limit")
	afterStr := c.Query("after")

	var limit *int

	if limitStr != "" {
		temp, err := strconv.Atoi(limitStr)

		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(api.ErrorResponse{
				Message: "Invalid Limit format",
				Code:    fiber.StatusBadRequest,
			})
			return
		}

		limit = &temp
	}

	var after *domain.NewsCursor

	if afterStr != "" {
		afterStrDecoded, err := base64.StdEncoding.DecodeString(afterStr)

		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(api.ErrorResponse{
				Message: "Invalid after format",
				Code:    fiber.StatusBadRequest,
			})
			return
		}

		value, err := strconv.ParseInt(string(afterStrDecoded), 10, 64)

		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(api.ErrorResponse{
				Message: "Invalid after format",
				Code:    fiber.StatusBadRequest,
			})
			return
		}

		after = &domain.NewsCursor{ID: value}
	}

	presenter := presenter.ListNewsPresenter{}
	inputDTO := dto.ListNewsInputDTO{
		Limit: limit,
		After: after,
	}

	repository, err := repository.NewRepository(&s.dbConfig)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(api.ErrorResponse{
			Message: fmt.Sprintf("Internal server error: %v", err),
			Code:    fiber.StatusInternalServerError,
		})
		return
	}

	useCase := usecase.ListNewsUseCase{
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
