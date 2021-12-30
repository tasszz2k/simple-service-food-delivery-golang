package restaurantstorage

import (
	"context"
	"simple-service-food-delivery-golang/common"
	"simple-service-food-delivery-golang/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := s.db

	if err := db.Where("id=?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
