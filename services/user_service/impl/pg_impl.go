package impl

import (
	"context"
	"drive/common"
	"drive/ent"
	"drive/ent/email"
	"drive/ent/social"
	"drive/ent/user"
	"drive/services/user_service"
	"fmt"
	"github.com/sirupsen/logrus"
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
	res, err = p.dbClient.User.Query().Where(user.DeletedAtEQ(common.ZeroTime), user.IDEQ(id)).First(ctx)

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
	query := p.dbClient.User.Query().Where(user.DeletedAtEQ(common.ZeroTime))
	if ids != nil {
		query = query.Where(user.IDIn(ids...))
	}
	res, err = query.All(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *pgImpl) DeleteOne(ctx context.Context, id int64) (res *ent.User, err error) {
	_user, err := p.dbClient.User.UpdateOneID(id).SetDeletedAt(time.Now()).Save(ctx)
	if err != nil {
		logrus.Errorf("soft delete user by id error: %v", err)
		return nil, err
	}
	return _user, nil
}

func (p *pgImpl) GetByPhone(ctx context.Context, phone string) (res *ent.User, err error) {
	_social, err := p.dbClient.Debug().Social.Query().Where(social.HasEmailWith(email.Name("hello"))).Count(ctx)
	fmt.Printf("%+v\n", _social)
	_user, err := p.dbClient.User.Query().Where(user.PhoneEQ(phone)).First(ctx)
	if err != nil {
		logrus.Errorf("get user by phone error: %v", err)
		return nil, err
	}
	return _user, nil
}
