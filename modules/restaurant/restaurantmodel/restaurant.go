package restaurantmodel

import (
	"errors"
	"simple-service-food-delivery-golang/common"
	"strings"
)

const EntityName = "restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	OwnerId         int    `json:"-" gorm:"column:owner_id"`
	Addr            string `json:"address" gorm:"column:addr"`
	CityId          int    `json:"city_id" gorm:"column:city_id"`
	LikeCount       int    `json:"like_count" gorm:"-"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name   *string `json:"name" gorm:"column:name"`
	Addr   *string `json:"address" gorm:"column:addr"`
	CityId int     `json:"city_id" gorm:"column:city_id"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	OwnerId         int    `json:"-" gorm:"column:owner_id"`
	Addr            string `json:"address" gorm:"column:addr"`
	CityId          int    `json:"city_id" gorm:"column:city_id"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name is required")
	}

	return nil
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "restaurant name cannot be empty", "ERR_NAME_CANNOT_BE_EMPTY")
)

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)
}
