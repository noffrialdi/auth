package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type SignUpRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone_name"`
	Address   string `json:"address"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (c SignUpRequest) Validation() error {
	err := validation.ValidateStruct(
		&c,
		validation.Field(&c.FirstName, validation.Required),
		validation.Field(&c.Phone, validation.Required),
		validation.Field(&c.Username, validation.Required),
		validation.Field(&c.Password, validation.Required),
	)
	if err != nil {
		return err
	}
	return err
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c SignInRequest) Validation() error {
	err := validation.ValidateStruct(
		&c,
		validation.Field(&c.Username, validation.Required),
		validation.Field(&c.Password, validation.Required),
	)
	if err != nil {
		return err
	}
	return err
}
