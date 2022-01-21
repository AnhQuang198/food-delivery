package service

import (
	"context"
	"food-delivery/commons"
	"food-delivery/modules/user/model"
)

type UserStore interface {
	Create(ctx context.Context, data *model.UserCreate) error
	GetUserByCondition(ctx context.Context, conditions map[string]interface{}) (*model.User, error)
}

type userService struct {
	store UserStore
}

func NewUserService(store UserStore) *userService {
	return &userService{store: store}
}

//register
func (biz *userService) Register(ctx context.Context, data *model.UserCreate) error {
	user, err := biz.store.GetUserByCondition(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return commons.ErrEntityExisted(model.EntityName, err)
	}

	//add validate data

	data.Role = "user"
	data.Password, err = commons.HashPassword(data.Password)
	if err != nil {
		return commons.ErrInternal(err)
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return commons.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}

//login
func (biz *userService) Login(ctx context.Context, data *model.UserLogin) (*model.Account, error) {
	user, err := biz.store.GetUserByCondition(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, model.ErrUsernameOrPasswordInvalid
	}

	if !commons.CheckPasswordHash(user.Password, data.Password) {
		return nil, model.ErrUsernameOrPasswordInvalid
	}

	return nil, nil

}
