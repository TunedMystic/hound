package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/matryer/is"

	"github.com/tunedmystic/hound/app/config"
	"github.com/tunedmystic/hound/app/database"
)

func init() {
	// Switch to project dir.
	config := config.GetConfig()
	os.Chdir(config.BaseDir)
}

func newTestServer() *Server {
	s := Server{}
	s.Templates = s.GetTemplates()
	s.DB = &sqlx.DB{}
	return &s
}

func Test_IndexHandler(t *testing.T) {
	is := is.New(t)

	// Mock the 'GetArticles' function.
	database.PerformGetArticles = func(db *sqlx.DB) []*database.Article {
		return []*database.Article{
			{URL: "http://example.com/article1", Title: "Article 1"},
			{URL: "http://example.com/article2", Title: "Article 2"},
			{URL: "http://example.com/article3", Title: "Article 3"},
		}
	}

	s := newTestServer()
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(s.IndexHandler).ServeHTTP(w, r)

	is.Equal(w.Code, http.StatusOK) // Status OK
}

func Test_VersionHandler(t *testing.T) {
	is := is.New(t)

	s := newTestServer()
	r := httptest.NewRequest("GET", "/version", nil)
	w := httptest.NewRecorder()

	http.HandlerFunc(s.VersionHandler).ServeHTTP(w, r)

	is.Equal(w.Code, http.StatusOK) // Status OK
}

func Test_NewServer(t *testing.T) {
	is := is.New(t)

	s := NewServer()

	is.True(s.DB != nil)        // DB exists
	is.True(s.Router != nil)    // Router exists
	is.True(s.Templates != nil) // Templates exist

	s.DB.Close()

	// Remove the test db.
	config := config.GetConfig()
	if err := os.Remove(config.DatabaseName); err != nil {
		t.Fatalf("Error when removing the database: %v\n", err)
	}
}
