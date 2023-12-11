package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ernazar151020/go-packages/config"
	"github.com/ernazar151020/go-packages/handlers"
	"github.com/ernazar151020/go-packages/pkg/render"
)

const port = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	app.UseCache = false

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	server := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	fmt.Println("Server is listens on port : ", port)

	err = server.ListenAndServe()
	log.Fatal(err)

	// http.ListenAndServe(port, nil)

}
