package v1

import "net/http"

type API struct {
	handler http.Handler
}

func NewAPI() (http.Handler, error) {
	api := &API{}
	api.registerRoutes()
	return api, nil
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.handler.ServeHTTP(w, r)
}
