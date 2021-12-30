package restaurantbiz

import (
	"context"
	"errors"
	"simple-service-food-delivery-golang/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{
		store: store,
	}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {

	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return err
	}

	if data.Status == 0 {
		return errors.New("restaurant has been deleted")
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return err
	}
	return nil
}
