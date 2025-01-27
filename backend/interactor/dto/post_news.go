package dto

type PostNewsInputDTO struct {
	ID          int64
	Title       *string
	Content     *string
	CategoryIDs []int64
}

type PostNewsOutputDTO struct {
}
