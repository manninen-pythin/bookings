package handlers

import (
	"net/http"

	"github.com/manninen-pythin/bookings/pkg/render"
)

// Home page handler, sends a writer and the name of the request to RenderTemplate function
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About page handler, sends a writer and the name of the request to RenderTemplate function
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
