package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-golang-04/common"
	"simple-service-golang-04/component"
	"simple-service-golang-04/modules/restaurant/restaurantbiz"
	"simple-service-golang-04/modules/restaurant/restaurantstorage"
)

func GetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		id := uid.GetLocalID()

		store := restaurantstorage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), int(id))

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
