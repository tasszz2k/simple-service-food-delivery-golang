package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-golang-04/common"
	"simple-service-golang-04/component"
	"simple-service-golang-04/modules/restaurant/restaurantbiz"
	"simple-service-golang-04/modules/restaurant/restaurantstorage"
	"strconv"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
