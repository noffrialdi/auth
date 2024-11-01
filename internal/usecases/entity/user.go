package entity

type SignUpResponse struct {
	Data    SignUpDataRespone
	Message string `json:"message"`
}

type SignUpDataRespone struct {
	Username string `json:"username"`
}

type SignInResponse struct {
	Message string `json:"message"`
}
