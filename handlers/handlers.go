// package handlers

// import (
// 	"net/http"

// 	"github.com/ernazar151020/go-packages/config"
// 	"github.com/ernazar151020/go-packages/models"
// 	"github.com/ernazar151020/go-packages/pkg/render"
// )

// var Repo *Repository

// type Repository struct {
// 	App *config.AppConfig
// }

// func NewRepo(a *config.AppConfig) *Repository {
// 	return &Repository{
// 		App: a,
// 	}
// }

// func NewHandler(r *Repository) {
// 	Repo = r
// }

// func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
// 	stringMap := make(map[string]string)
// 	stringMap["test"] = "Ernazar"

// 	remoteIP := req.RemoteAddr
// 	m.App.Session.Put(req.Context(), "remote_ip", remoteIP)

// 	render.RenderTemplate(res, "home.pages.html", &models.TemplateData{
// 		StringMap: stringMap,
// 	})
// }

// func (m *Repository) About(res http.ResponseWriter, req *http.Request) {

// 	stringMap := make(map[string]string)
// 	stringMap["test"] = "Ernazar"

// 	remoteIP := m.App.Session.GetString(req.Context(), "remote_ip")

// 	stringMap["remote_ip"] = remoteIP

// 	render.RenderTemplate(res, "about.pages.html", &models.TemplateData{StringMap: stringMap})

// }

package handlers

import (
	"net/http"

	"github.com/ernazar151020/go-packages/config"
	"github.com/ernazar151020/go-packages/models"
	"github.com/ernazar151020/go-packages/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Ernazar"

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "home.pages.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.pages.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
