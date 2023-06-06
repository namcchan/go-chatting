package usecase

import (
	"context"
	"github.com/namcchan/go-chatting/internal/domain"
	"github.com/namcchan/go-chatting/internal/repository"
	"github.com/namcchan/go-chatting/pkg/utils"
)

type authUseCase struct {
	ctx context.Context
}

func NewAuthUseCase(ctx context.Context) domain.AuthUseCase {
	return &authUseCase{
		ctx: ctx,
	}
}

func (a *authUseCase) Register(payload *domain.RegisterPayload) error {
	password, err := utils.HashPassword(payload.Password)
	if err != nil {
		return err
	}
	userRepository := repository.NewUserRepository(a.ctx)

	err = userRepository.Create(&domain.User{
		Username: payload.Username,
		Email:    payload.Email,
		Name:     payload.Name,
		Password: password,
	})

	if err != nil {
		return err
	}

	return nil
}

func (a *authUseCase) Login(payload *domain.LoginPayload) (*domain.TokenData, error) {
	//TODO implement me
	panic("implement me")
}

func (a *authUseCase) ForgotPassword() error {
	//TODO implement me
	panic("implement me")
}

func (a *authUseCase) ResetPassword(resetPasswordToken string, password string) error {
	//TODO implement me
	panic("implement me")
}

func (a *authUseCase) GetMe() (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
