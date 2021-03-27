package handlers

import (
	"github.com/WilliamAinsworth/bookings/pkg/config"
	"github.com/WilliamAinsworth/bookings/pkg/models"
	"github.com/WilliamAinsworth/bookings/pkg/render"
	"net/http"
)

// Repo the repository
var Repo *Respoitory

// Repository pattern type
type Respoitory struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Respoitory {
	return &Respoitory{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Respoitory) {
	Repo = r
}

// Home is the home page handler
func (m *Respoitory) Home(w http.ResponseWriter, r *http.Request) {
	// everytime user hits the home page, their IP is stored in the session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Respoitory) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl",  &models.TemplateData{
		StringMap: stringMap,
	})
}
