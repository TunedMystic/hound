package database

import "time"

// Article represents a news article.
type Article struct {
	ID          int       `db:"id"`
	URL         string    `db:"url" json:"url"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	PublishedAt time.Time `db:"published_at" json:"publishedAt"`
	CreatedAt   time.Time `db:"created_at"`
}
