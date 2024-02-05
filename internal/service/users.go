package service

import (
	"context"
	"fmt"

	"github.com/slaughtlaught/web-service-api/internal/entity"
	"github.com/slaughtlaught/web-service-api/internal/value"
	"github.com/slaughtlaught/web-service-api/pkg/errorx"
	"golang.org/x/crypto/bcrypt"
)

type userSource interface {
	SaveUser(context.Context, entity.User) error
	GetUserByID(context.Context, value.UserID) (entity.User, error)
	GetUserByEmail(context.Context, value.Email) (entity.User, error)
}

type Users struct {
	userSource userSource
}

func NewUsers(userSource userSource) Users {
	return Users{
		userSource: userSource,
	}
}

func (u Users) GetByID(ctx context.Context, id value.UserID) (entity.User, error) {
	user, err := u.userSource.GetUserByID(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("userSource.GetUserByID: %w", err)
	}

	return user, nil
}

func (u Users) GetByEmail(ctx context.Context, email value.Email) (entity.User, error) {
	user, err := u.userSource.GetUserByEmail(ctx, email)
	if err != nil {
		return entity.User{}, fmt.Errorf("userSource.GetUserByEmail: %w", err)
	}

	return user, nil
}

func (u Users) Save(ctx context.Context, user entity.User, password string) error {
	user, err := u.GetByEmail(ctx, user.Email)
	if err != nil && !errorx.IsNotFoundError(err) {
		return fmt.Errorf("u.GetByEmail: %w", err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}

	user.HashedPassword = string(hashedPassword)

	if err := u.userSource.SaveUser(ctx, user); err != nil {
		return fmt.Errorf("userSource.SaveUser: %w", err)
	}
	return nil
}
