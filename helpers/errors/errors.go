package errors

import "errors"

var (
	ERR_EMAIL_IS_EMPTY           = errors.New("Email is empty")
	ERR_NAME_IS_EMPTY            = errors.New("Name is empty")
	ERR_PASSWORD_IS_EMPTY        = errors.New("Password is empty")
	ERR_GET_ADMIN_BAD_REQUEST_ID = errors.New("Can't find any admin with this id")
	ERR_GET_USER_BAD_REQUEST_ID  = errors.New("Can't find any user with this id")
	ERR_DELETE_ADMIN             = errors.New("Can't delete this admin")
	ERR_DELETE_USER              = errors.New("Can't delete this user")
	ERR_DELETE_ADMIN             = errors.New("Can't delete this admin")
	ERR_REGISTER_ADMIN_DATABASE  = errors.New("Failed to register new admin to database")
	ERR_BCRYPT_PASSWORD          = errors.New("Failed to bcrypt password")
	ERR_UPDATE_DATA              = errors.New("Failed to update data")
	ERR_TOKEN                    = errors.New("Failed to create new token")
	ERR_WRONG_PASSWORD           = errors.New("Wrong password")
	ERR_EMAIL_NOT_FOUND          = errors.New("Email not found")
	ERR_REGISTER_USER_DATABASE   = errors.New("Failed to register new user to database")
	ERR_GET_DATA                 = errors.New("Database can't request any data right now")

