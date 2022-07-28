package directory_service

import (
	"context"
	"drive/ent"
)

type Repository interface {
	Get(ctx context.Context, id int64) (res *ent.Directory, err error)
	Create(ctx context.Context, name string, parentId int64, userId int64, isPublic bool) (res *ent.Directory, err error)
	List(ctx context.Context, ids []int64) (res ent.Directories, err error)
	DeleteOne(ctx context.Context, id int64, userId int64) (res *ent.Directory, err error)
	UpdateOne(ctx context.Context, id int64, attr map[string]interface{}, userId int64) (res *ent.Directory, err error)
	ListByParentId(ctx context.Context, parentId int64, userId int64) (res ent.Directories, err error)
}
