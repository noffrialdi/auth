package dao

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/noffrialdi/auth/internal/model"
)

func TestUserRepo_Insert(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxMock := sqlx.NewDb(db, "sqlmock")
	mock.ExpectBegin()
	txMock, err := sqlxMock.Beginx()
	if err != nil {
		t.Errorf("tx error: %v", err)
		return
	}

	ctx := context.Background()
	Id := uuid.New()
	// Now := time.Now()

	payload := model.User{
		Id: Id,
	}

	r := NewUserRepo(&OptsUser{
		DB: sqlxMock,
	})

	type args struct {
		ctx  context.Context
		dbTx *sqlx.Tx
		req  *model.User
	}
	tests := []struct {
		name    string
		repo    *UserRepo
		args    args
		want    uuid.UUID
		wantErr bool
		fn      func()
	}{
		{
			name: "Success with dbtx",
			args: args{
				ctx:  ctx,
				req:  &payload,
				dbTx: txMock,
			},
			want:    Id,
			wantErr: false,
			fn: func() {
				mock.ExpectExec(regexp.QuoteMeta(insertUser)).WithArgs(
					payload.Id,
					payload.FirstName,
					payload.LastName,
					payload.PhoneNumber,
					payload.Address,
					payload.Username,
					payload.Password,
					payload.UserIDCreated,
					payload.CreatedTime,
					payload.UserIDUpdated,
					payload.UpdatedTime,
					payload.DeletedTime,
					payload.IsDeleted,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
		},
		{
			name: "Failed error connection",
			args: args{
				ctx:  ctx,
				req:  &payload,
				dbTx: txMock,
			},
			wantErr: true,
			fn: func() {
				mock.ExpectExec(regexp.QuoteMeta(insertUser)).WillReturnError(errors.New("Failed"))
			},
		},
	}
	for _, tt := range tests {
		tt.fn()
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.Insert(tt.args.ctx, tt.args.dbTx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthRepo.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepo.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepo_GetByUserName(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxMock := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()
	r := NewUserRepo(&OptsUser{
		DB: sqlxMock,
	})

	query := getByUsernameQuery + ` where username = ?`

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
		fn      func()
	}{
		{
			name: "Success",
			args: args{
				ctx:      ctx,
				username: "test",
			},

			wantErr: false,
			want: &model.User{
				Username: "test",
			},
			fn: func() {
				columns := sqlmock.NewRows([]string{
					"username",
				}).AddRow("test")
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("test").WillReturnRows(columns)
			},
		},
		{
			name: "Failed error connection",
			args: args{
				ctx:      ctx,
				username: "test",
			},
			wantErr: true,
			fn: func() {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(sql.ErrConnDone)
			},
		},
	}
	for _, tt := range tests {
		tt.fn()
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetByUserName(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthRepo.GetByUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthRepo.GetByUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
