package server

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/tunedmystic/hound/app/database"
)

// Server struct ...
type Server struct {
	Router    *mux.Router
	Templates *template.Template
	DB        *sqlx.DB
}

func NewServer() *Server {
	s := Server{}
	s.DB = s.GetDatabase()
	s.Router = s.GetRouter()
	s.Templates = s.GetTemplates()
	return &s
}

func (s *Server) GetDatabase() *sqlx.DB {
	db := database.NewDB()
	database.CreateTables(db)
	return db
}

func (s *Server) GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", s.IndexHandler).Methods("GET")
	router.HandleFunc("/version", s.VersionHandler)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return router
}

func (s *Server) GetTemplates() *template.Template {
	templatePath := "templates/*.html"
	tmpl := template.Must(template.New("").ParseGlob(templatePath))
	return tmpl
}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {
	articles := database.GetArticles(s.DB)

	// Template context
	data := struct {
		Articles []*database.Article
	}{
		Articles: articles,
	}

	s.Templates.ExecuteTemplate(w, "index", data)
}

func (s *Server) VersionHandler(w http.ResponseWriter, r *http.Request) {
	s.Templates.ExecuteTemplate(w, "version", nil)
}
