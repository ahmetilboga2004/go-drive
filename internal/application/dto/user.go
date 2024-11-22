package dto

type UserRegister struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
	Username  string `json:"username" validate:"required,min=6,max=20,alphanum"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type UserLogin struct {
	UsernameOrEmail string `json:"username_or_email" validate:"required,min=6,max=20"`
	Password        string `json:"password" validate:"required,min=8"`
}

type UserBasicInfo struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}

type UserDetails struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email,omitempty"`
}
