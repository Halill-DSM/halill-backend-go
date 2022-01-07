package repository

import (
	"context"
	"halill/ent"
	"halill/ent/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserRepository interface {
	GetByEmail(string) (*ent.User, error)
	CreateUser(*ent.User) (*ent.User, error)
}

type userRepositoryImpl struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (ur *userRepositoryImpl) GetByEmail(email string) (*ent.User, error) {
	u, err := ur.db.User.Query().
		Where(user.ID(email)).
		Only(context.Background())
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "존재하지 않는 사용자 입니다.")
		}
		return nil, err
	}

	return u, nil
}

func (ur *userRepositoryImpl) CreateUser(user *ent.User) (*ent.User, error) {
	u, err := ur.db.User.Create().
		SetID(user.ID).
		SetPassword(user.Password).
		SetName(user.Name).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return u, nil
}
