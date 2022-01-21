package repository

import (
	"context"
	"food-delivery/commons"
	"food-delivery/modules/user/model"
)

func (s *sqlStore) Create(ctx context.Context, data *model.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(model.UserCreate{}.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return commons.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return commons.ErrDB(err)
	}

	return nil
}
