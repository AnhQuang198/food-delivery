package repository

import (
	"context"
	"food-delivery/commons"
	"food-delivery/modules/restaurant/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *model.RestaurantUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return commons.ErrDB(err)
	}

	return nil

}
