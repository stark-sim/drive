package impl

import (
	"context"
	"drive/ent"
	"drive/services/user"
)

type pgImpl struct {
	dbClient *ent.Client
}

func NewPgImpl(dbClient *ent.Client) user.Repository {
	return &pgImpl{
		dbClient: dbClient,
	}
}

func (p *pgImpl) Get(id int64) (user ent.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (p *pgImpl) Create(ctx context.Context, name string, password string, phone string) (user ent.User, err error) {
	user = *p.dbClient.User.Create().SetName(name).SetPassword(password).SetPhone(phone).SaveX(ctx)

	return user, nil
}

func (p *pgImpl) List(ids []int64) (users []ent.User, err error) {
	//TODO implement me
	panic("implement me")
}
