package domain

type RegisterPayload struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthUseCase interface {
	Register(payload *RegisterPayload) error
	Login(payload *LoginPayload) (*TokenData, error)
	ForgotPassword() error
	ResetPassword(resetPasswordToken string, password string) error
	GetMe() (*User, error)
}
