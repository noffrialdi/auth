package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/noffrialdi/auth/internal/usecases/entity"
	"github.com/noffrialdi/auth/internal/usecases/interactor/mocks_interactor"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAPI_Signup(t *testing.T) {
	mockUserInteractor := &mocks_interactor.UserIteractor{}

	req := `{
			"first_name" :"test",
			"username" : "test",
			"phone_name" : "09288199331",
			"address" : "test",
			"password" : "test"
		}`
	opts := &Opts{
		DefaultTimeout: 1,
		User:           mockUserInteractor,
	}

	server := httprouter.New()
	server = New(opts).Register()

	type args struct {
		body io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		fn         func()
	}{
		{
			name: "Error empty payload",
			args: args{
				body: nil,
			},
			wantStatus: 400,
			fn: func() {
			},
		},

		{
			name: "Success",
			args: args{
				body: strings.NewReader(req),
			},
			wantStatus: 200,
			fn: func() {
				mockUserInteractor.On("Signup", mock.Anything, mock.Anything).Return(&entity.SignUpResponse{}, http.StatusOK, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.fn()
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/user/signup", tt.args.body)
			resp := httptest.NewRecorder()

			server.ServeHTTP(resp, req)
			require.Equalf(t, tt.wantStatus, resp.Code, "Want status '%d', got '%d'", tt.wantStatus, resp.Code)
		})
	}
}

func TestAPI_Signin(t *testing.T) {
	mockUserInteractor := &mocks_interactor.UserIteractor{}

	req := `{
			"username" : "test",
			"password" : "test"
		}`
	opts := &Opts{
		DefaultTimeout: 1,
		User:           mockUserInteractor,
	}

	server := httprouter.New()
	server = New(opts).Register()

	type args struct {
		body io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		fn         func()
	}{
		{
			name: "Error empty payload",
			args: args{
				body: nil,
			},
			wantStatus: 400,
			fn: func() {
			},
		},

		{
			name: "Success",
			args: args{
				body: strings.NewReader(req),
			},
			wantStatus: 200,
			fn: func() {
				mockUserInteractor.On("Signin", mock.Anything, mock.Anything).Return(&entity.SignInResponse{}, http.StatusOK, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.fn()
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/user/signin", tt.args.body)
			resp := httptest.NewRecorder()

			server.ServeHTTP(resp, req)
			require.Equalf(t, tt.wantStatus, resp.Code, "Want status '%d', got '%d'", tt.wantStatus, resp.Code)
		})
	}
}
