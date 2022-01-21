package service

import (
	"context"
	"errors"
	"food-delivery/commons"
	"food-delivery/modules/restaurant/model"
)

type RestaurantStore interface {
	Create(ctx context.Context, data *model.RestaurantCreate) error
	Update(ctx context.Context, id int, data *model.RestaurantUpdate) error
	GetListByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *model.Filter,
		paging *commons.Paging,
	) ([]model.Restaurant, error)
	GetByCondition(
		ctx context.Context,
		conditions map[string]interface{},
	) (*model.Restaurant, error)
	SoftDelete(ctx context.Context, id int) error
}

type restaurantService struct {
	store RestaurantStore
}

func NewRestaurantService(store RestaurantStore) *restaurantService {
	return &restaurantService{store: store}
}

//create Restaurant
func (restaurant *restaurantService) CreateRestaurant(ctx context.Context, data *model.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	return restaurant.store.Create(ctx, data)
}

//update Restaurant
func (restaurant *restaurantService) UpdateRestaurant(ctx context.Context, id int, data *model.RestaurantUpdate) error {
	oldData, err := restaurant.store.GetByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := restaurant.store.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}

//get Restaurant by condition
func (restaurant *restaurantService) GetListRestaurants(ctx context.Context,
	conditions map[string]interface{},
	filter *model.Filter,
	paging *commons.Paging,
) ([]model.Restaurant, error) {
	results, err := restaurant.store.GetListByCondition(ctx, nil, filter, paging)
	return results, err
}

//get one
func (restaurant *restaurantService) GetRestaurant(
	ctx context.Context,
	id int,
) (*model.Restaurant, error) {
	data, err := restaurant.store.GetByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != commons.RecordNotFound {
			return nil, commons.ErrCannotGetEntity(model.EntityName, err)
		}
		return nil, err
	}

	if data.Status == 0 {
		return nil, errors.New("data deleted")
	}

	return data, nil
}

//delete soft
func (restaurant *restaurantService) Delete(ctx context.Context, id int) error {
	oldData, err := restaurant.store.GetByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := restaurant.store.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}
