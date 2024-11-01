package dao

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/noffrialdi/auth/internal/domain/repository"
	"github.com/noffrialdi/auth/internal/model"
	"go.opentelemetry.io/otel"
)

type OptsUser struct {
	DB *sqlx.DB
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(o *OptsUser) repository.UserRepository {
	return &UserRepo{
		db: o.DB,
	}
}

const insertUser = `INSERT INTO user
	(
		id,
		first_name,
		last_name,
		phone_number,
		address,
		username,
		password,
		user_id_created,
		created_time,
		user_id_updated,
		updated_time,
		deleted_time,
		is_deleted
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ;`

func (repo *UserRepo) Insert(ctx context.Context, dbTx *sqlx.Tx, req *model.User) (uuid.UUID, error) {

	var idUuid uuid.UUID
	if req.Id != uuid.Nil {
		idUuid = req.Id
	} else {
		idUuid = uuid.New()
	}

	args := []interface{}{
		idUuid,
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
		req.Address,
		req.Username,
		req.Password,
		req.UserIDCreated,
		req.CreatedTime,
		req.UserIDUpdated,
		req.UpdatedTime,
		req.DeletedTime,
		req.IsDeleted,
	}

	var err error
	if dbTx == nil {
		_, err = repo.db.ExecContext(ctx, insertUser, args...)
	} else {
		_, err = dbTx.ExecContext(ctx, insertUser, args...)
	}

	if err != nil {
		return uuid.Nil, err
	}

	return idUuid, nil
}

const getByUsernameQuery = `SELECT id, first_name, last_name, phone_number, username,address, password, user_id_created, user_id_updated, created_time, updated_time, deleted_time, is_deleted from user`

func (repo *UserRepo) GetByUserName(ctx context.Context, username string) (*model.User, error) {
	ctx, span := otel.Tracer("test").Start(ctx, "AuthRepository.GetByUserName")
	defer span.End()
	var (
		data model.User
		err  error
	)

	query := getByUsernameQuery + ` where username = ?`
	err = repo.db.GetContext(ctx, &data, query, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return &data, nil
}
