package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-acme/lego/log"
	"github.com/gophish/hub/config"
)

type Server struct {
	config       *config.ServerConfig
	http         *http.Server
	repositories []string
}

type ServerOption func(*Server) error

func NewServer(serverConfig *config.ServerConfig, opts ...ServerOption) (*Server, error) {
	defaultServer := &http.Server{
		ReadTimeout: 10 * time.Second,
		Addr:        serverConfig.ListenAddress,
	}
	server := &Server{
		config: serverConfig,
		http:   defaultServer,
	}
	for _, opt := range opts {
		if err := opt(server); err != nil {
			return server, err
		}
	}
	err := server.registerRoutes()
	return server, err
}

func (server *Server) Start() error {
	log.Infof("Starting healthcheck server at http://%s", server.config.ListenAddress)
	return server.http.ListenAndServe()
}

// Shutdown attempts to gracefully shutdown the server.
func (server *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return server.http.Shutdown(ctx)
}
