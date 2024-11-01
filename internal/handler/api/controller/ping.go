package controller

import (
	"net/http"

	"github.com/noffrialdi/auth/internal/infrastructures/custerr"
)

func (a *API) Ping(w http.ResponseWriter, r *http.Request) {
	s := "Pong"
	var data interface{}
	custerr.RespondWithSuccess(w, s, data)
}
