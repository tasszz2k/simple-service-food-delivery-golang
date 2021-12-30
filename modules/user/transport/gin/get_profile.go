package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-food-delivery-golang/common"
	"simple-service-food-delivery-golang/component"
)

func GetProfile(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)
		//data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
