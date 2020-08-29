package database

import (
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/matryer/is"
	_ "github.com/mattn/go-sqlite3" // sqlite

	"github.com/tunedmystic/hound/app/config"
)

func createArticles(db *sqlx.DB) func() {
	db.MustExec(`
		insert into article (url, title, description, published_at, created_at)
		values
		("cnn.com/article1", "Article 1 title", "Article 1 description", "2020-07-17 07:38:44", "2020-07-20 07:38:44"),
		("cnn.com/article2", "Article 2 title", "Article 2 description", "2020-07-01 01:06:31", "2020-07-01 04:00:21"),
		("msnbc.com/article2", "Article 2 title", "Article 2 description", "2020-06-12 13:21:00", "2020-06-13 21:00:00");
	`)

	return func() {
		// Truncate table and reset PK sequence.
		db.MustExec(`delete from article; delete from sqlite_sequence where name='article';`)
	}
}

func TestMain(m *testing.M) {
	// Switch to project dir.
	config := config.GetConfig()
	os.Chdir(config.BaseDir)

	// Create test db, and create tables.
	db := NewDB()
	CreateTables(db)
	if err := db.Close(); err != nil {
		log.Fatal("Error when closing the database")
	}

	// Run the test.
	code := m.Run()

	// Remove the test db.
	os.Remove(config.DatabaseName)
	os.Exit(code)
}

func Test_Database(t *testing.T) {
	is := is.New(t)

	db := NewDB()
	cleanup := createArticles(db)
	defer cleanup()

	articles := GetArticles(db)

	is.Equal(len(articles), 3) // 3 Articles exist
}
