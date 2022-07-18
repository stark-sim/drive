package impl

import (
	"context"
	"drive/ent"
	"drive/ent/user"
	"drive/services/user_service"
	"time"
)

type pgImpl struct {
	dbClient *ent.Client
}

func NewPgImpl(dbClient *ent.Client) user_service.Repository {
	return &pgImpl{
		dbClient: dbClient,
	}
}

func (p *pgImpl) Get(ctx context.Context, id int64) (res *ent.User, err error) {
	res, err = p.dbClient.User.Query().Where(user.DeletedAtEQ(time.Time{}), user.IDEQ(id)).First(ctx)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) Create(ctx context.Context, name string, password string, phone string) (res *ent.User, err error) {
	res, err = p.dbClient.User.Create().SetName(name).SetPassword(password).SetPhone(phone).Save(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pgImpl) List(ctx context.Context, ids []int64) (res ent.Users, err error) {
	query := p.dbClient.User.Query().Where(user.DeletedAtEQ(time.Time{}))
	if ids != nil {
		query = query.Where(user.IDIn(ids...))
	}
	res, err = query.All(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
