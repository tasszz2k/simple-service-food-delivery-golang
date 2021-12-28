package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-golang-04/common"
	"simple-service-golang-04/component"
	"simple-service-golang-04/modules/restaurant/restaurantbiz"
	"simple-service-golang-04/modules/restaurant/restaurantmodel"
	"simple-service-golang-04/modules/restaurant/restaurantstorage"
)

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
