package user_service

import (
	"context"
	"drive/ent"
)

type Repository interface {
	Get(ctx context.Context, id int64) (res *ent.User, err error)
	Create(ctx context.Context, name string, password string, phone string) (res *ent.User, err error)
	List(ctx context.Context, ids []int64) (res ent.Users, err error)
}
