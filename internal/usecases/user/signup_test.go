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
	"github.com/noffrialdi/auth/internal/usecases/entity"
	"github.com/stretchr/testify/mock"
)

func Test_module_Signup(t *testing.T) {
	mockTx := &mocks_interface.TxManager{}
	mockUser := &mocks_domain.UserRepository{}
	id := uuid.New()
	type args struct {
		ctx context.Context
		req *request.SignUpRequest
	}
	tests := []struct {
		name           string
		args           args
		wantRes        *entity.SignUpResponse
		wantStatusCode int
		wantErr        bool
		fn             func()
	}{
		{
			name: "ShouldSuccess",
			args: args{
				ctx: context.Background(),
				req: &request.SignUpRequest{
					FirstName: "test",
					LastName:  "test",
					Phone:     "test",
					Address:   "test",
					Username:  "test",
					Password:  "test",
				},
			},
			wantRes: &entity.SignUpResponse{
				Message: "Signup berhasil",
				Data: entity.SignUpDataRespone{
					Username: "test",
				},
			},
			wantErr:        false,
			wantStatusCode: http.StatusOK,
			fn: func() {
				mockUser.Mock.On("GetByUserName", mock.Anything, mock.Anything).Return(nil, errors.New("data tidak ditemukan")).Once()
				mockUser.Mock.On("Insert", mock.Anything, mock.Anything, mock.Anything).Return(id, nil).Once()
			},
		},
		{
			name: "ShouldError_GetByUserName",
			args: args{
				ctx: context.Background(),
				req: &request.SignUpRequest{
					FirstName: "test",
					LastName:  "test",
					Phone:     "test",
					Address:   "test",
					Username:  "test",
					Password:  "test",
				},
			},

			wantErr:        true,
			wantStatusCode: http.StatusBadRequest,
			fn: func() {
				mockUser.Mock.On("GetByUserName", mock.Anything, mock.Anything).Return(nil, errors.New("Failed")).Once()
			},
		},
		{
			name: "ShouldError_Insert",
			args: args{
				ctx: context.Background(),
				req: &request.SignUpRequest{
					FirstName: "test",
					LastName:  "test",
					Phone:     "test",
					Address:   "test",
					Username:  "test",
					Password:  "test",
				},
			},

			wantErr:        true,
			wantStatusCode: http.StatusBadRequest,
			fn: func() {
				mockUser.Mock.On("GetByUserName", mock.Anything, mock.Anything).Return(nil, errors.New("data tidak ditemukan")).Once()
				mockUser.Mock.On("Insert", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("Failed")).Once()
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
			gotRes, gotStatusCode, err := m.Signup(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("module.Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("module.Signup() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("module.Signup() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}
