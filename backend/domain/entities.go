package domain

type News struct {
	ID       int64   `json:"Id"`
	Title    string  `json:"Title"`
	Content  string  `json:"Content"`
	Category []int64 `json:"Category"`
}

type NewsCursor struct {
	ID int64
}
