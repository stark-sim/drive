package impl

import (
	"context"
	"drive/common"
	"drive/ent"
	"drive/ent/directory"
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
	res, err = p.dbClient.Directory.Query().Where(directory.DeletedAtEQ(common.ZeroTime), directory.IDEQ(id)).First(ctx)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) Create(ctx context.Context, name string, parentId int64, userId int64, isPublic bool) (res *ent.Directory, err error) {
	res, err = p.dbClient.Directory.Create().
		SetName(name).
		SetParentID(parentId).
		SetCreatedBy(userId).
		SetUpdatedBy(userId).
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

func (p *pgImpl) DeleteOne(ctx context.Context, id int64, userId int64) (res *ent.Directory, err error) {
	res, err = p.dbClient.Directory.UpdateOneID(id).SetDeletedAt(time.Now()).SetUpdatedBy(userId).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) UpdateOne(ctx context.Context, id int64, attr map[string]interface{}, userId int64) (res *ent.Directory, err error) {
	res, err = p.dbClient.Directory.UpdateOneID(id).
		SetUpdatedBy(userId).
		SetName(attr["name"].(string)).
		SetParentID(attr["parentId"].(int64)).
		SetIsPublic(attr["isPublic"].(bool)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) ListByParentId(ctx context.Context, parentId int64, userId int64) (res ent.Directories, err error) {
	query := p.dbClient.Directory.Query()
	if parentId == 0 {
		query = query.Where(
			directory.And(
				directory.Not(directory.HasParent()),
				directory.Or(
					directory.IsPublicEQ(true),
					directory.CreatedBy(userId),
				),
			),
		)
	} else {
		query = query.Where(
			directory.And(
				directory.ParentIDEQ(parentId),
				directory.Or(
					directory.IsPublicEQ(true),
					directory.CreatedBy(userId),
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
