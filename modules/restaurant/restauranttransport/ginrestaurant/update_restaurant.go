package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-golang-04/common"
	"simple-service-golang-04/component"
	"simple-service-golang-04/modules/restaurant/restaurantbiz"
	"simple-service-golang-04/modules/restaurant/restaurantmodel"
	"simple-service-golang-04/modules/restaurant/restaurantstorage"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			return
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
