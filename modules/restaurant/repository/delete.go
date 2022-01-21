package repository

import (
	"context"
	"food-delivery/commons"
	"food-delivery/modules/restaurant/model"
)

func (s *sqlStore) SoftDelete(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(model.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return commons.ErrDB(err)
	}
	return nil
}
