package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/noffrialdi/auth/internal/handler/api/request"
	"github.com/noffrialdi/auth/internal/infrastructures/custerr"
)

func (a *API) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var payload request.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		custerr.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = payload.Validation()
	if err != nil {
		custerr.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, statusCode, err := a.User.Signup(ctx, &payload)
	if err != nil {
		custerr.RespondWithError(w, statusCode, err.Error())
		return
	}

	custerr.RespondWithSuccess(w, "Sukses", res)
}

func (a *API) Signin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var payload request.SignInRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		custerr.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = payload.Validation()
	if err != nil {
		custerr.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, statusCode, err := a.User.Signin(ctx, &payload)
	if err != nil {
		custerr.RespondWithError(w, statusCode, err.Error())
		return
	}

	custerr.RespondWithSuccess(w, "Sukses", res)
}
