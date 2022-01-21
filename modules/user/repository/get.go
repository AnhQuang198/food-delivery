package repository

import (
	"context"
	"food-delivery/commons"
	"food-delivery/modules/user/model"
	"gorm.io/gorm"
)

func (s *sqlStore) GetUserByCondition(ctx context.Context, conditions map[string]interface{}) (*model.User, error) {
	db := s.db.Table(model.User{}.TableName())
	var user model.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, commons.RecordNotFound
		}
		return nil, commons.ErrDB(err)
	}

	return &user, nil
}
