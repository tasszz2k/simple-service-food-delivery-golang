package userbiz

import (
	"context"
	"simple-service-food-delivery-golang/common"
	"simple-service-food-delivery-golang/component"
	"simple-service-food-delivery-golang/component/tokenprovider"
	usermodel "simple-service-food-delivery-golang/modules/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

//type TokenConfig interface {
//	GetAtExp() int
//	GetRtExp() int
//}

type loginBusiness struct {
	appCtx        component.AppContext
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher,
	expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry, // TODO: should convert to TokenConfig with at_exp and rt_exp
	}
}

// 1. Find user, email
// 2. Hash password from input and compare with password from db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s) to client

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	passHashed := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordIncorrect
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//refreshToken, err := business.tokenProvider.Generate(payload, business.tkCfg.GetRtExp())
	//if err != nil {
	//	return nil, common.ErrInternal(err)
	//}
	//
	//account := usermodel.NewAccount(accessToken, refreshToken)

	return accessToken, nil
}
