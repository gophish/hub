package server

import (
	"time"

	"github.com/NYTimes/gziphandler"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gophish/hub/server/api/v1"
)

func (server *Server) registerRoutes() error {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(gziphandler.GzipHandler)
	r.Use(middleware.Timeout(60 * time.Second))

	v1API, err := v1.NewAPI()
	if err != nil {
		return err
	}
	r.Mount("/api/v1", v1API)
	server.http.Handler = r
	return nil
}
