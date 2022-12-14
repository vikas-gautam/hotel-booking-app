package handlers

import (
	"net/http"

	"github.com/vikas-gautam/hotel-booking-app/pkg/config"
	"github.com/vikas-gautam/hotel-booking-app/pkg/models"
	"github.com/vikas-gautam/hotel-booking-app/pkg/render"
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
	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	stringMap["remote_ip"] = remoteIP

	//send the data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
