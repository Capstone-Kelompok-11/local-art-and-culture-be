package errors

import (
	"net/http"
)

func GetCodeError(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
