package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/noffrialdi/auth/internal/model"
)

type UserRepository interface {
	Insert(ctx context.Context, dbTx *sqlx.Tx, req *model.User) (uuid.UUID, error)
	GetByUserName(ctx context.Context, username string) (*model.User, error)
}
