package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-food-delivery-golang/common"
	"simple-service-food-delivery-golang/component"
	"simple-service-food-delivery-golang/modules/restaurant/restaurantbiz"
	"simple-service-food-delivery-golang/modules/restaurant/restaurantmodel"
	"simple-service-food-delivery-golang/modules/restaurant/restaurantstorage"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			return
		}

		id := uid.GetLocalID()

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			return
		}

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), int(id), &data); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
