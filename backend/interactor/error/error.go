package prjerror

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "record not found"
}
