package api

type PostNewsRequest struct {
	Title       *string
	Content     *string
	CategoryIDs []int64
}

type PostNewsResponse struct {
	Success bool
}
