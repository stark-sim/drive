package object_service

import (
	"context"
	"drive/ent"
)

type Repository interface {
	Get(ctx context.Context, id int64) (res *ent.Object, err error)
	Create(ctx context.Context, name string, url string, directoryID int64, userID int64, isPublic bool) (res *ent.Object, err error)
	List(ctx context.Context, ids []int64) (res ent.Objects, err error)
	DeleteOne(ctx context.Context, id int64, userID int64) (res *ent.Object, err error)
	UpdateOne(ctx context.Context, id int64, attr map[string]interface{}, userID int64) (res *ent.Object, err error)
	ListByDirectoryID(ctx context.Context, directoryID int64, userID int64) (res ent.Objects, err error)
}
