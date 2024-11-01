package auth

import (
	"github.com/noffrialdi/auth/internal/domain/repository"
	"github.com/noffrialdi/auth/internal/interfaces/txmanager"
	"github.com/noffrialdi/auth/internal/usecases/interactor"
)

type module struct {
	txManager txmanager.TxManager
	userRepo  repository.UserRepository
}

type Opts struct {
	TxManager txmanager.TxManager
	UserRepo  repository.UserRepository
}

func New(o *Opts) interactor.UserIteractor {
	return &module{
		txManager: o.TxManager,
		userRepo:  o.UserRepo,
	}
}
