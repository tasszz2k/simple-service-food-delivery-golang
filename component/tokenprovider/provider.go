package tokenprovider

import (
	"errors"
	"simple-service-food-delivery-golang/common"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ERROR_TOKEN_NOT_FOUND",
	)
	ErrInvalidToken = common.NewCustomError(
		errors.New("invalid token"),
		"invalid token provided",
		"ERROR_INVALID_TOKEN",
	)

	ErrEncodingToken = common.NewCustomError(
		errors.New("error encoding token"),
		"error encoding token",
		"ERROR_ENCODING_TOKEN",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}
