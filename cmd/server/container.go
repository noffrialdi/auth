package server

import (
	"github.com/jmoiron/sqlx"
	"github.com/noffrialdi/auth/config"
	"github.com/noffrialdi/auth/internal/interfaces/dao"
	"github.com/noffrialdi/auth/internal/interfaces/txmanager"
	"github.com/noffrialdi/auth/internal/usecases/interactor"
	auth "github.com/noffrialdi/auth/internal/usecases/user"
)

type Container struct {
	Cfg  config.MainConfig
	Auth interactor.UserIteractor
}

type Opts struct {
	Cfg          *config.MainConfig
	MasterDataDB *sqlx.DB
}

func newContainer(o *Opts) *Container {
	tx := txmanager.NewTxManager(&txmanager.Opts{
		DB: o.MasterDataDB,
	})

	userRepo := dao.NewUserRepo(&dao.OptsUser{
		DB: o.MasterDataDB,
	})

	authInteractor := auth.New(&auth.Opts{
		TxManager: tx,
		UserRepo:  userRepo,
	})

	return &Container{
		Cfg:  *o.Cfg,
		Auth: authInteractor,
	}
}
