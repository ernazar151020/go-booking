package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ernazar151020/go-packages/internal/config"
	"github.com/ernazar151020/go-packages/internal/handlers"
	"github.com/ernazar151020/go-packages/internal/render"
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
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	app.UseCache = false

	render.NewTemplates(&app)

	server := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	fmt.Println("Server is listens on port : ", port)

	err = server.ListenAndServe()
	log.Fatal(err)

}
