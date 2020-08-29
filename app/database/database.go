package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // sqlite

	"github.com/tunedmystic/hound/app/config"
)

// NewDB ...
func NewDB() *sqlx.DB {
	config := config.GetConfig()
	db := sqlx.MustOpen("sqlite3", config.DatabaseName)
	return db
}

// CreateTables ...
func CreateTables(db *sqlx.DB) {
	PerformCreateTables(db)
}

// GetArticles ...
func GetArticles(db *sqlx.DB) []*Article {
	return PerformGetArticles(db)
}

// PerformCreateTables [internal]
var PerformCreateTables = func(db *sqlx.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS article (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			url VARCHAR(100) UNIQUE NOT NULL,
			title VARCHAR(100) NOT NULL,
			description VARCHAR(100),
			published_at DATETIME NOT NULL,
			created_at DATETIME NOT NULL
		);`

	db.MustExec(sql)
}

// PerformGetArticles [internal]
var PerformGetArticles = func(db *sqlx.DB) []*Article {
	articles := []*Article{}
	sql := `SELECT * FROM article;`

	if err := db.Select(&articles, sql); err != nil {
		fmt.Printf("Could not fetch articles: %v\n", err.Error())
	}

	return articles
}
