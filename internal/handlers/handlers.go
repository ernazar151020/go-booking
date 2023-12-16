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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ernazar151020/go-packages/internal/config"
	"github.com/ernazar151020/go-packages/internal/forms"
	"github.com/ernazar151020/go-packages/internal/models"
	"github.com/ernazar151020/go-packages/internal/render"
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

	render.RenderTemplate(w, r, "home.pages.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.pages.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "general.pages.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
func (m *Repository) MajorsSuit(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "majors.pages.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "make-reservation.pages.html", &models.TemplateData{
		Forms: forms.New(nil),
	})

}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

}
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "search-availability.pages.html", &models.TemplateData{
		StringMap: stringMap,
	})

}
func (m *Repository) PosetSearch(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start date is %s, end date is %s", start, end)))
}

type responseData struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {

	response := responseData{
		OK:      true,
		Message: "Available",
	}

	res, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling")
	}

	fmt.Println(string(res))

	w.Write(res)

}
