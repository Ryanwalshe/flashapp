package handlers

import (
	"net/http"

	"github.com/Ryanwalshe/flashcard_app/pkg/config"
	"github.com/Ryanwalshe/flashcard_app/pkg/render"
)

// the respository used by the handlers
var Repo *Respository

type Respository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Respository {
	return &Respository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Respository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
