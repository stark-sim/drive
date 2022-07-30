package directory_service

import (
	"context"
	"drive/ent"
)

type Repository interface {
	Get(ctx context.Context, id int64) (res *ent.Directory, err error)
	Create(ctx context.Context, name string, parentID int64, userID int64, isPublic bool) (res *ent.Directory, err error)
	List(ctx context.Context, ids []int64) (res ent.Directories, err error)
	DeleteOne(ctx context.Context, id int64, userID int64) (res *ent.Directory, err error)
	UpdateOne(ctx context.Context, id int64, attr map[string]interface{}, userID int64) (res *ent.Directory, err error)
	ListByParentID(ctx context.Context, parentID int64, userID int64) (res ent.Directories, err error)
}
