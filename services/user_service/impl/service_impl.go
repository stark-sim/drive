package impl

import (
	"context"
	"drive/services/user_service"
	"drive/tools"
	"errors"
	"github.com/sirupsen/logrus"
	"time"
)

type UserServiceImpl struct {
	repository user_service.Repository
}

func NewUserService(repository user_service.Repository) user_service.UserService {
	return &UserServiceImpl{repository: repository}
}

func (u *UserServiceImpl) Login(ctx context.Context, phone, password string) (token string, err error) {
	user, err := u.repository.GetByPhone(ctx, phone)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		logrus.Infof("password incorrect: %v != %v", user.Password, password)
		return "", errors.New("password incorrect")
	}

	token, err = tools.GetToken(time.Now(), user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
