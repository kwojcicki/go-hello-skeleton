package rest

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	// TODO: added liveness/readiness probe
	//r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Use(middleware.StripSlashes)
	options := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "X-Auth-Token", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
	cors := cors.New(options)
	r.Use(cors.Handler)
	r.Use(middleware.NoCache)
	r.Use(middleware.CloseNotify)


	r.Route("/", func(r chi.Router) {
		r.Route("/pet", func(r chi.Router) {
			r.Get("/{petID}", GetPet)
		})
		// TODO: rest of CRUD endpoints
		//r.Route("/docs", func(r chi.Router) {
		//	r.Get("/", GetDocs)
		//	r.Get("/{docId}", GetDocByID)
		//	r.Put("/", SaveDoc)
		//})
	})

	return r
}