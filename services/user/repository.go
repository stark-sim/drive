package user

import (
	"context"
	"drive/ent"
)

type Repository interface {
	Get(id int64) (user ent.User, err error)
	Create(ctx context.Context, name string, password string, phone string) (user ent.User, err error)
	List(ids []int64) (users []ent.User, err error)
}
