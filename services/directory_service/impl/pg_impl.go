package impl

import (
	"context"
	"drive/common"
	"drive/ent"
	"drive/ent/directory"
	"drive/ent/object"
	"drive/services/directory_service"
	"time"
)

type pgImpl struct {
	dbClient *ent.Client
}

func NewPgImpl(dbClient *ent.Client) directory_service.Repository {
	return &pgImpl{
		dbClient: dbClient,
	}
}

func (p *pgImpl) Get(ctx context.Context, id int64) (res *ent.Directory, err error) {
	res, err = p.dbClient.Directory.Query().
		Where(directory.DeletedAtEQ(common.ZeroTime), directory.IDEQ(id)).
		WithChildren(func(query *ent.DirectoryQuery) {
			query.Where(directory.DeletedAtEQ(common.ZeroTime)).Order(ent.Asc(directory.FieldCreatedAt))
		}).
		WithObjects(func(query *ent.ObjectQuery) {
			query.Where(object.DeletedAtEQ(common.ZeroTime)).Order(ent.Asc(object.FieldCreatedAt))
		}).
		First(ctx)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) Create(ctx context.Context, name string, parentID int64, userID int64, isPublic bool) (res *ent.Directory, err error) {
	res, err = p.dbClient.Directory.Create().
		SetName(name).
		SetParentID(parentID).
		SetCreatedBy(userID).
		SetUpdatedBy(userID).
		SetIsPublic(isPublic).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) List(ctx context.Context, ids []int64) (res ent.Directories, err error) {
	res, err = p.dbClient.Directory.Query().Where(directory.DeletedAtEQ(common.ZeroTime), directory.IDIn(ids...)).
		Order(ent.Asc(directory.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) DeleteOne(ctx context.Context, id int64, userID int64) (res *ent.Directory, err error) {
	res, err = p.dbClient.Directory.UpdateOneID(id).SetDeletedAt(time.Now()).SetUpdatedBy(userID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) UpdateOne(ctx context.Context, id int64, attr map[string]interface{}, userID int64) (res *ent.Directory, err error) {
	res, err = p.dbClient.Directory.UpdateOneID(id).
		SetUpdatedBy(userID).
		SetName(attr["name"].(string)).
		SetParentID(attr["parentID"].(int64)).
		SetIsPublic(attr["isPublic"].(bool)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) ListByParentID(ctx context.Context, parentID int64, userID int64) (res ent.Directories, err error) {
	query := p.dbClient.Directory.Query()
	if parentID == 0 {
		query = query.Where(
			directory.And(
				directory.Not(directory.HasParent()),
				directory.Or(
					directory.IsPublicEQ(true),
					directory.CreatedBy(userID),
				),
			),
		)
	} else {
		query = query.Where(
			directory.And(
				directory.ParentIDEQ(parentID),
				directory.Or(
					directory.IsPublicEQ(true),
					directory.CreatedBy(userID),
				),
			),
		)
	}

	res, err = query.Order(ent.Asc(directory.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
