package impl

import (
	"context"
	"drive/ent"
	"drive/services/directory_service"
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
	//TODO implement me
	panic("implement me")
}

func (p *pgImpl) Create(ctx context.Context, name string, password string, phone string) (res *ent.Directory, err error) {
	//TODO implement me
	panic("implement me")
}

func (p *pgImpl) List(ctx context.Context, ids []int64) (res ent.Directories, err error) {
	//TODO implement me
	panic("implement me")
}

func (p *pgImpl) DeleteOne(ctx context.Context, id int64) (res *ent.Directory, err error) {
	//TODO implement me
	panic("implement me")
}

func (p *pgImpl) UpdateOne(ctx context.Context, id int64, attr map[string]interface{}) (res *ent.Directory, err error) {
	//TODO implement me
	panic("implement me")
}
