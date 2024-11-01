package handler

import (
	"log"
	"net"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/noffrialdi/auth/config"
	"github.com/noffrialdi/auth/internal/handler/api/controller"
	"github.com/noffrialdi/auth/internal/usecases/interactor"
	graceful "gopkg.in/tylerb/graceful.v1"
)

type Opts struct {
	Cfg  config.MainConfig
	User interactor.UserIteractor
}

type Handler struct {
	options     *Opts
	listenErrCh chan error
	router      *httprouter.Router
}

func NewHTTP(o *Opts) *Handler {
	handler := &Handler{options: o}
	handler.router = controller.New(&controller.Opts{
		User: o.User,
	}).Register()

	return handler
}

func (h *Handler) Run() {
	log.Printf("API Listening on %s", h.options.Cfg.Server.Port)

	l, err := net.Listen("tcp4", h.options.Cfg.Server.Port)
	if err != nil {
		log.Panicf("Error tcp4 %s", err)
	}

	h.listenErrCh <- graceful.Serve(&http.Server{
		Addr:         h.options.Cfg.Server.Port,
		ReadTimeout:  h.options.Cfg.Server.ReadTimeout,
		WriteTimeout: h.options.Cfg.Server.WriteTimeout,
		Handler:      h.router,
	}, l, h.options.Cfg.Server.GracefulTimeout)
}

func (h *Handler) ListenError() <-chan error {
	return h.listenErrCh
}
