package restaurantstorage

import (
	"context"
	"simple-service-golang-04/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {
	db := s.db

	if err := db.Where("id=?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
