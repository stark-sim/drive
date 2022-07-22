package user_service

import "context"

type UserService interface {
	Login(ctx context.Context, phone, password string) (token string, err error)
}
