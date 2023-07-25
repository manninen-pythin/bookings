package handlers

import (
	"net/http"

	"github.com/manninen-pythin/bookings/pkg/config"
	"github.com/manninen-pythin/bookings/pkg/render"
)

// The repository used by NewHandlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler, sends a writer and the name of the request to RenderTemplate function
func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About page handler, sends a writer and the name of the request to RenderTemplate function
func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
