package restaurantbiz

import (
	"context"
	"errors"
	"simple-service-food-delivery-golang/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	UpdateData(
		ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{
		store: store,
	}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(
	ctx context.Context,
	id int,
	newData *restaurantmodel.RestaurantUpdate) error {

	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return err
	}

	if data.Status == 0 {
		return errors.New("restaurant has been deleted")
	}

	if err := biz.store.UpdateData(ctx, id, newData); err != nil {
		return err
	}
	return nil
}
