package interactor

import (
	"context"

	"github.com/noffrialdi/auth/internal/handler/api/request"
	"github.com/noffrialdi/auth/internal/usecases/entity"
)

type UserIteractor interface {
	Signup(ctx context.Context, req *request.SignUpRequest) (respone *entity.SignUpResponse, statusCode int, err error)
	Signin(ctx context.Context, req *request.SignInRequest) (respone *entity.SignInResponse, statusCode int, err error)
}
