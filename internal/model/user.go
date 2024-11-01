package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id            uuid.UUID      `db:"id"`
	FirstName     string         `db:"first_name"`
	LastName      sql.NullString `db:"last_name"`
	PhoneNumber   sql.NullString `db:"phone_number"`
	Username      string         `db:"username"`
	Address       sql.NullString `db:"address"`
	Password      string         `db:"password"`
	UserIDCreated int            `db:"user_id_created"`
	CreatedTime   time.Time      `db:"created_time"`
	UserIDUpdated int            `db:"user_id_updated"`
	UpdatedTime   time.Time      `db:"updated_time"`
	DeletedTime   sql.NullTime   `db:"deleted_time"`
	IsDeleted     int8           `db:"is_deleted"`
}
