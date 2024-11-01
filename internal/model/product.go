package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id            uuid.UUID `db:"id"`
	ProductName   string    `db:"product_name"`
	ProductDesc   string    `db:"product_desc"`
	UserIDCreated int       `db:"user_id_created"`
	CreatedTime   time.Time `db:"created_time"`
	UserIDUpdated int       `db:"user_id_updated"`
	UpdatedTime   time.Time `db:"updated_time"`
	IsDeleted     int8      `db:"is_deleted"`
}
