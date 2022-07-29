package impl

import (
	"context"
	"drive/common"
	"drive/ent"
	"drive/ent/directory"
	"drive/ent/object"
	"drive/ent/user"
	"drive/services/object_service"
	"time"
)

type pgImpl struct {
	dbClient *ent.Client
}

func NewPgImpl(dbClient *ent.Client) object_service.Repository {
	return &pgImpl{
		dbClient: dbClient,
	}
}

func (p *pgImpl) Get(ctx context.Context, id int64) (res *ent.Object, err error) {
	res, err = p.dbClient.Object.Query().Where(object.DeletedAtEQ(common.ZeroTime), object.IDEQ(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (p *pgImpl) Create(ctx context.Context, name string, url string, directoryID int64, userID int64, isPublic bool) (res *ent.Object, err error) {
	res, err = p.dbClient.Object.Create().
		SetName(name).
		SetURL(url).
		SetDirectoryID(directoryID).
		SetUserID(userID).
		SetCreatedBy(userID).
		SetUpdatedBy(userID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) List(ctx context.Context, ids []int64) (res ent.Objects, err error) {
	res, err = p.dbClient.Object.Query().Where(object.IDIn(ids...), object.DeletedAtEQ(common.ZeroTime)).Order(ent.Asc(object.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) DeleteOne(ctx context.Context, id int64, userID int64) (res *ent.Object, err error) {
	res, err = p.dbClient.Object.UpdateOneID(id).SetUpdatedBy(userID).SetDeletedAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) UpdateOne(ctx context.Context, id int64, attr map[string]interface{}, userID int64) (res *ent.Object, err error) {
	res, err = p.dbClient.Object.UpdateOneID(id).SetUpdatedBy(userID).SetName(attr["name"].(string)).SetIsPublic(attr["isPublic"].(bool)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) ListByDirectoryID(ctx context.Context, directoryID int64, userID int64) (res ent.Objects, err error) {
	res, err = p.dbClient.Object.Query().
		Where(
			object.HasDirectoryWith(directory.IDEQ(directoryID)),
			object.Or(object.IsPublicEQ(true), object.HasUserWith(user.IDEQ(userID))),
		).
		Order(ent.Asc(object.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
