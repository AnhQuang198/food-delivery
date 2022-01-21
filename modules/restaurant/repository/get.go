package repository

import (
	"context"
	"food-delivery/commons"
	"food-delivery/modules/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) GetListByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	filter *model.Filter,
	paging *commons.Paging,
) ([]model.Restaurant, error) {
	var results []model.Restaurant

	db := s.db
	db = db.Table(model.Restaurant{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if len(v.Name) > 0 {
			db = db.Where("name = ?", v.Name)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, commons.ErrDB(err)
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit). //so dong se bo qua khi lay record
		Limit(paging.Limit).                      //gioi han so dong lay ra
		Order("id desc").
		Where("status in (1)").
		Find(&results).Error; err != nil {
		return nil, commons.ErrDB(err)
	}

	return results, nil
}

func (s *sqlStore) GetByCondition(
	ctx context.Context,
	conditions map[string]interface{},
) (*model.Restaurant, error) {
	var data model.Restaurant

	db := s.db
	if err := db.Where(conditions).Where("status in (1)").First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, commons.RecordNotFound
		}
		return nil, commons.ErrDB(err)
	}

	return &data, nil
}
