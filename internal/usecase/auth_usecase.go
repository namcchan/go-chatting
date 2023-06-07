package usecase

import (
	"context"
	"errors"
	"github.com/namcchan/go-chatting/internal/domain"
	"github.com/namcchan/go-chatting/internal/repository"
	"github.com/namcchan/go-chatting/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type authUseCase struct {
	ctx context.Context
}

func NewAuthUseCase(ctx context.Context) domain.AuthUseCase {
	return &authUseCase{
		ctx: ctx,
	}
}

func (ac *authUseCase) Register(payload *domain.RegisterPayload) error {
	password, err := utils.HashPassword(payload.Password)
	if err != nil {
		return err
	}
	userRepository := repository.NewUserRepository(ac.ctx)

	filter := bson.M{
		"$or": bson.A{
			bson.M{"username": payload.Username},
			bson.M{"email": payload.Email},
		},
	}
	existedUser, err := userRepository.FindOne(filter)
	if err != nil {
		return err
	}

	if existedUser != nil {
		return errors.New("username or email already in used")
	}

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

func (ac *authUseCase) Login(payload *domain.LoginPayload) (*domain.Tokens, error) {
	userRepository := repository.NewUserRepository(ac.ctx)

	filter := bson.M{"username": payload.Username}

	user, err := userRepository.FindOne(filter)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("username or password is incorrect")
	}

	isCorrectPassword := utils.CheckPasswordHash(payload.Password, user.Password)
	if !isCorrectPassword {
		return nil, errors.New("incorrect password")
	}

	accessExpires := time.Now().Add(time.Minute * 15).Unix()
	refreshExpires := time.Now().Add(time.Minute * 24 * 30).Unix()

	accessToken, refreshToken, err := utils.GenerateTokens(user.ID, accessExpires, refreshExpires)
	if err != nil {
		return nil, errors.New("error when generate tokens")
	}

	tokens := &domain.Tokens{
		Access: domain.TokenPayload{
			Value:   accessToken,
			Expires: accessExpires,
		},
		Refresh: domain.TokenPayload{
			Value:   refreshToken,
			Expires: refreshExpires,
		},
	}
	return tokens, nil
}

func (ac *authUseCase) ForgotPassword() error {
	//TODO implement me
	panic("implement me")
}

func (ac *authUseCase) ResetPassword(resetPasswordToken string, password string) error {
	//TODO implement me
	panic("implement me")
}

func (ac *authUseCase) GetMe() (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
