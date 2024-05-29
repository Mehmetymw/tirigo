package dtos

type UserRegisterParameter struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserAuthParameter struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
