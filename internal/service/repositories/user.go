package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/ArtemGretsov/teletodo/internal/service/entities"
)

type User interface {
	CreateAuth(ctx context.Context, user entities.User) (entities.User, error)
	GetByAuth(ctx context.Context, userAuthentication entities.UserAuthentication) (entities.User, error)
	GetAll(ctx context.Context) (users entities.Users, err error)
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return &user{db: db}
}

func (u *user) CreateAuth(ctx context.Context, userEntity entities.User) (entities.User, error) {
	err := u.db.WithContext(ctx).
		Where("auth_type = ?", userEntity.UserAuthentication.AuthType).
		Where("external_id = ?", userEntity.UserAuthentication.ExternalID).
		Find(&userEntity.UserAuthentication).
		Error

	if err != nil {
		return userEntity, errors.WithStack(err)
	}

	if userEntity.UserAuthentication.UserID == 0 {
		userEntity.UUID = uuid.New().String()
		err = u.db.WithContext(ctx).
			Session(&gorm.Session{AllowGlobalUpdate: true}).
			Create(&userEntity).Error
		if err != nil {
			return userEntity, errors.WithStack(err)
		}

		return userEntity, nil
	}

	userEntity.ID = userEntity.UserAuthentication.UserID
	err = u.db.WithContext(ctx).Select("Name").Save(&userEntity).Error
	if err != nil {
		return userEntity, errors.WithStack(err)
	}

	return userEntity, nil
}

func (u *user) GetByAuth(ctx context.Context, userAuth entities.UserAuthentication) (userEntity entities.User, err error) {
	err = u.db.WithContext(ctx).
		First(&userAuth, &userAuth).
		Error
	if err != nil {
		return userEntity, errors.WithStack(err)
	}

	err = u.db.First(&userEntity, userAuth.UserID).Error
	if err != nil {
		return userEntity, errors.WithStack(err)
	}

	userEntity.UserAuthentication = &userAuth

	return userEntity, nil
}

func (u *user) GetAll(ctx context.Context) (users entities.Users, err error) {
	err = u.db.WithContext(ctx).Preload("UserAuthentication").Find(&users).Error

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return
}
