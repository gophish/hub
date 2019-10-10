package v1

import "net/http"

func (api *API) Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
