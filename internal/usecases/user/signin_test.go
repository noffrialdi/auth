package auth

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/noffrialdi/auth/internal/domain/mocks_domain"
	"github.com/noffrialdi/auth/internal/handler/api/request"
	"github.com/noffrialdi/auth/internal/interfaces/mocks_interface"
	"github.com/noffrialdi/auth/internal/model"
	"github.com/noffrialdi/auth/internal/usecases/entity"
	"github.com/stretchr/testify/mock"
)

func Test_module_Signin(t *testing.T) {
	mockTx := &mocks_interface.TxManager{}
	mockUser := &mocks_domain.UserRepository{}
	id := uuid.New()

	type args struct {
		ctx context.Context
		req *request.SignInRequest
	}
	tests := []struct {
		name           string
		args           args
		wantRes        *entity.SignInResponse
		wantStatusCode int
		wantErr        bool
		fn             func()
	}{
		{
			name: "ShouldSuccess",
			args: args{
				ctx: context.Background(),
				req: &request.SignInRequest{
					Username: "test",
					Password: "test",
				},
			},
			wantStatusCode: http.StatusOK,
			wantRes: &entity.SignInResponse{
				Message: "Sukses login",
			},
			wantErr: false,
			fn: func() {
				mockUser.Mock.On("GetByUserName", mock.Anything, mock.Anything).Return(&model.User{
					Id:       id,
					Username: "test",
					Password: "$2a$10$Et7viOQvlsS9GeC9RK7MHuZBLkCZPNDsib90VI0idd315soMbxDK2",
				}, nil).Once()
			},
		},
		{
			name: "ShouldError_password_not_match",
			args: args{
				ctx: context.Background(),
				req: &request.SignInRequest{
					Username: "test",
					Password: "test",
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
			fn: func() {
				mockUser.Mock.On("GetByUserName", mock.Anything, mock.Anything).Return(&model.User{
					Id:       id,
					Username: "test",
					Password: "test",
				}, nil).Once()
			},
		},
		{
			name: "ShouldError_data_not_found",
			args: args{
				ctx: context.Background(),
				req: &request.SignInRequest{
					Username: "test",
					Password: "test",
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
			fn: func() {
				mockUser.Mock.On("GetByUserName", mock.Anything, mock.Anything).Return(nil, errors.New("data tidak ditemukan")).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.fn()
		t.Run(tt.name, func(t *testing.T) {
			m := &module{
				txManager: mockTx,
				userRepo:  mockUser,
			}
			gotRes, gotStatusCode, err := m.Signin(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("module.Signin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("module.Signin() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("module.Signin() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}
