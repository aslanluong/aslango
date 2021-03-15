package main

import (
	"aslango/pkg/routes"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main()  {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/go", routes.LinkRoute)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Oops, nothing here :("))
	})

	http.ListenAndServe(":9999", r)
}