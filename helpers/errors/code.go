package errors

import (
	"net/http"
)

func GetCodeError(err error) int {
	switch err {
	case ERR_EMAIL_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_PASSWORD_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_NAME_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_DELETE_ADMIN:
		return http.StatusInternalServerError
	case ERR_DELETE_USER:
		return http.StatusInternalServerError
	case ERR_GET_ADMIN_BAD_REQUEST_ID:
		return http.StatusNotFound
	case ERR_GET_USER_BAD_REQUEST_ID:
		return http.StatusNotFound
	case ERR_GET_ADMIN_BAD_REQUEST_ID:
		return http.StatusNotFound
	case ERR_GET_DATA:
		return http.StatusInternalServerError
	case ERR_REGISTER_ADMIN_DATABASE:
		return http.StatusInternalServerError
	case ERR_REGISTER_USER_DATABASE:
		return http.StatusInternalServerError
	case ERR_BCRYPT_PASSWORD:
		return http.StatusInternalServerError
	case ERR_TOKEN:
		return http.StatusInternalServerError
	case ERR_WRONG_PASSWORD:
		return http.StatusBadRequest
	case ERR_EMAIL_NOT_FOUND:
		return http.StatusNotFound
	case ERR_UPDATE_DATA:
		return http.StatusInternalServerError
	case ERR_ROLE_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_GET_ROLE_BAD_REQUEST_ID:
		return http.StatusNotFound
	case ERR_DELETE_ROLE:
		return http.StatusInternalServerError
	case ERR_TITLE_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_CONTENT_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_PHONE_NUMBER_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_OUTLET_NAME_IS_EMPTY:
		return http.StatusBadRequest
	case ERR_CREATE_ROLE:
		return http.StatusInternalServerError
	case ERR_GET_CREATOR_BAD_REQUEST_ID:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
