package common

import (
	"errors"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnthorizedErrorResponse(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "Database error", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid request", err.Error(), "INVALID_REQUEST")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"Something went wrong in the server",
		err.Error(),
		"INTERNAL_ERROR")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		"Cannot list "+strings.ToLower(entity),
		"ERROR_CANNOT_LIST_"+strings.ToUpper(entity))
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		"Cannot delete "+strings.ToLower(entity),
		"ERROR_CANNOT_DELETE_"+strings.ToUpper(entity))
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		"Cannot get "+strings.ToLower(entity),
		"ERROR_CANNOT_GET_"+strings.ToUpper(entity))
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		"Entity "+strings.ToLower(entity)+" already existed",
		"ERROR_ENTITY_EXISTED_"+strings.ToUpper(entity))
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		"Entity "+strings.ToLower(entity)+" already deleted",
		"ERROR_ENTITY_DELETED_"+strings.ToUpper(entity))
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		"Entity "+strings.ToLower(entity)+" not found",
		"ERROR_ENTITY_NOT_FOUND_"+strings.ToUpper(entity))
}
