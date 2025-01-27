package db_entities

//go:generate reform

// NewsCategories represents a row in news_categories table.
//
//reform:news_categories
type NewsCategories struct {
	ID         int32 `reform:"id,pk"`
	NewsID     int64 `reform:"news_id"`
	CategoryID int64 `reform:"category_id"`
}
