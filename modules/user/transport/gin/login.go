package ginuser

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-service-food-delivery-golang/common"
	"simple-service-food-delivery-golang/component"
	"simple-service-food-delivery-golang/component/hasher"
	"simple-service-food-delivery-golang/component/tokenprovider/jwt"
	userbiz "simple-service-food-delivery-golang/modules/user/biz"
	usermodel "simple-service-food-delivery-golang/modules/user/model"
	userstorage "simple-service-food-delivery-golang/modules/user/storage"
)

func Login(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSqlStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}

}
