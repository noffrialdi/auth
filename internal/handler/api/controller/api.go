package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/noffrialdi/auth/internal/usecases/interactor"
)

type API struct {
	defaultTimeout int
	User           interactor.UserIteractor
}

type Opts struct {
	DefaultTimeout int
	User           interactor.UserIteractor
}

func New(o *Opts) *API {
	return &API{
		defaultTimeout: o.DefaultTimeout,
		User:           o.User,
	}
}

func (a *API) Register() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc("GET", "/health", a.Ping)

	router.HandlerFunc("POST", "/user/signup", a.Signup)
	router.HandlerFunc("POST", "/user/signin", a.Signin)
	return router
}
