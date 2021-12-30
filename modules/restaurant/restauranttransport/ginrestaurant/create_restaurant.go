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

func CreateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.OwnerId = requester.GetUserId()

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.GenUID(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
