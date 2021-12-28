package restaurantmodel

import (
	"errors"
	"simple-service-golang-04/common"
	"strings"
)

type Restaurant struct {
	common.SQLModel `json:,inline`
	Name            string `json:"name" gorm:"column:name"`
	Addr            string `json:"address" gorm:"column:addr"`
	CityId          int    `json:"city_id" gorm:"column:city_id"`
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
	Id     int    `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Addr   string `json:"address" gorm:"column:addr"`
	CityId int    `json:"city_id" gorm:"column:city_id"`
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
