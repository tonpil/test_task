package db_entities

//go:generate reform

// News represents a row in news table.
//
//reform:news
type News struct {
	ID      int64  `reform:"id,pk"`
	Title   string `reform:"title"`
	Content string `reform:"content"`
}
