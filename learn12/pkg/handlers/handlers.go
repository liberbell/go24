package handlers

import (
	"net/http"

	"github.com/tsawler/go-course/pkg/config"
	"github.com/tsawler/go-course/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	StringMap := make(map[string]string)
	StringMap["test"] = "Hello, again."
	render.RenderTemplate(w, "about.page.tmpl", &TemplateData{
		StringMap: StringMap,
	})
}
