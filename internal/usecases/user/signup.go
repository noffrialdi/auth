package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/noffrialdi/auth/internal/handler/api/request"
	"github.com/noffrialdi/auth/internal/infrastructures/utils"
	"github.com/noffrialdi/auth/internal/model"
	"github.com/noffrialdi/auth/internal/usecases/entity"
)

func (m *module) Signup(ctx context.Context, req *request.SignUpRequest) (res *entity.SignUpResponse, statusCode int, err error) {

	user, err := m.userRepo.GetByUserName(ctx, req.FirstName)
	if err != nil && err.Error() != "data tidak ditemukan" {
		log.Println("[GetByUserName] error GetByUserName ", err)
		return nil, http.StatusBadRequest, errors.New("failed")
	}

	if user != nil {
		return nil, http.StatusBadRequest, errors.New("username telah terdaftar")
	}

	now := time.Now()

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		log.Println("[HashPassword] error hash password ", err)
		return nil, http.StatusBadRequest, errors.New("error hash password")
	}

	_, err = m.userRepo.Insert(ctx, nil, &model.User{
		Username:    req.Username,
		FirstName:   req.FirstName,
		LastName:    sql.NullString{String: req.LastName},
		PhoneNumber: sql.NullString{String: req.Phone},
		Password:    hash,
		Address:     sql.NullString{String: req.Address},
		CreatedTime: now,
		UpdatedTime: now,
	})
	if err != nil {
		log.Println("[Insert] error Insert user ", err)
		return nil, http.StatusBadRequest, err
	}

	result := entity.SignUpResponse{
		Message: "Signup berhasil",
		Data: entity.SignUpDataRespone{
			Username: req.Username,
		},
	}
	return &result, http.StatusOK, nil
}
