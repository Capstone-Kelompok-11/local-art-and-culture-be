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
	case ERR_GET_ADMIN_BAD_REQUEST_ID:
		return http.StatusBadRequest
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
	default:
		return http.StatusInternalServerError
	}
}
