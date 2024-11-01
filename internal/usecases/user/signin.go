package auth

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/noffrialdi/auth/internal/handler/api/request"
	"github.com/noffrialdi/auth/internal/infrastructures/utils"
	"github.com/noffrialdi/auth/internal/usecases/entity"
)

func (m *module) Signin(ctx context.Context, req *request.SignInRequest) (res *entity.SignInResponse, statusCode int, err error) {

	user, err := m.userRepo.GetByUserName(ctx, req.Username)
	if err != nil {
		if err.Error() == "data tidak ditemukan" {
			return nil, http.StatusBadRequest, errors.New("data tidak ditemukan")
		}
		log.Println("[GetByUserName] error GetByUserName ", err)
		return nil, http.StatusBadRequest, errors.New("error GetByUserName")
	}

	match := utils.CheckPassword(req.Password, user.Password)
	if !match {
		return nil, http.StatusBadRequest, errors.New("password tidak sesuai, silahkan coba lagi")
	}
	result := entity.SignInResponse{
		Message: "Sukses login",
	}

	return &result, http.StatusOK, nil
}
