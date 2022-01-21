package repository

import (
	"context"
	"food-delivery/commons"
	"food-delivery/modules/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *model.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return commons.ErrDB(err)
	}

	return nil
}
