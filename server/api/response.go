package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Success  bool        `json:"success"`
	Error    *string     `json:"error"`
	Response interface{} `json:"response"`
}

func JSONError(w http.ResponseWriter, r *http.Request, err error, c int) {
	errText := err.Error()
	response := &Response{
		Error:   &errText,
		Success: false,
	}
	render.Status(r, c)
	render.JSON(w, r, response)
}

func JSONResponse(w http.ResponseWriter, r *http.Request, data interface{}, c int) {
	response := &Response{
		Success:  true,
		Response: data,
	}
	render.Status(r, c)
	render.JSON(w, r, response)
}
