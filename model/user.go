package model

import (
	"context"
	"database/sql"

	"github.com/gami/simple_arch_example/gen/schema"
	"github.com/gami/simple_arch_example/mysql"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type User interface {
	FindByID(ctx context.Context, id uint64) (user *schema.User, err error)
	Create(ctx context.Context) (id uint64, err error)
}

type user struct {
	db *sql.DB
}

func NewUser(db *mysql.DB) User {
	return &user{
		db: db.DB,
	}
}

func (m *user) FindByID(ctx context.Context, id uint64) (*schema.User, error) {
	u, err := schema.Users(
		schema.UserWhere.ID.EQ(id),
	).One(ctx, m.db)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to query user id=%v", id)
	}
	return u, nil
}

func (m *user) Create(ctx context.Context) (uint64, error) {
	u := &schema.User{}
	err := u.Insert(ctx, m.db, boil.Infer())
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert user")
	}
	return u.ID, nil
}
