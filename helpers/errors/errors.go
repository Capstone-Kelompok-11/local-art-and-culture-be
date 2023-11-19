package errors

import "errors"

var (
	ERR_EMAIL_IS_EMPTY              = errors.New("email is empty")
	ERR_PRICE_IS_EMPTY              = errors.New("price is empty")
	ERR_DESCRIPTION_IS_EMPTY        = errors.New("description is empty")
	ERR_STATUS_IS_EMPTY             = errors.New("status is empty")
	ERR_NAME_IS_EMPTY               = errors.New("name is empty")
	ERR_OUTLET_NAME_IS_EMPTY        = errors.New("outlet name is empty")
	ERR_CATEGORY_IS_EMPTY           = errors.New("category is empty")
	ERR_TYPE_IS_EMPTY               = errors.New("type is empty")
	ERR_PASSWORD_IS_EMPTY           = errors.New("password is empty")
	ERR_PHONE_NUMBER_IS_EMPTY       = errors.New("phone number is empty")
	ERR_GET_ADMIN_BAD_REQUEST_ID    = errors.New("can't find any admin with this id")
	ERR_GET_USER_BAD_REQUEST_ID     = errors.New("can't find any user with this id")
	ERR_DELETE_ADMIN                = errors.New("can't delete this admin")
	ERR_DELETE_USER                 = errors.New("can't delete this user")
	ERR_REGISTER_ADMIN_DATABASE     = errors.New("failed to register new admin to database")
	ERR_BCRYPT_PASSWORD             = errors.New("failed to bcrypt password")
	ERR_UPDATE_DATA                 = errors.New("failed to update data")
	ERR_TOKEN                       = errors.New("failed to create new token")
	ERR_WRONG_PASSWORD              = errors.New("wrong password")
	ERR_EMAIL_NOT_FOUND             = errors.New("email not found")
	ERR_REGISTER_USER_DATABASE      = errors.New("failed to register new user to database")
	ERR_GET_DATA                    = errors.New("database can't request any data right now")
	ERR_ROLE_IS_EMPTY               = errors.New("role is empty")
	ERR_GET_ROLE_BAD_REQUEST_ID     = errors.New("can't find any role with this id")
	ERR_GET_CREATOR_BAD_REQUEST_ID  = errors.New("can't find any creator with this id")
	ERR_GET_PRODUCT_BAD_REQUEST_ID  = errors.New("can't find any product with this id")
	ERR_GET_CATEGORY_BAD_REQUEST_ID = errors.New("can't find any category with this id")
	ERR_DELETE_ROLE                 = errors.New("can't delete this role")
	ERR_DELETE_CATEGORY             = errors.New("can't delete this category")
	ERR_TITLE_IS_EMPTY              = errors.New("title is empty")
	ERR_CONTENT_IS_EMPTY            = errors.New("content is empty")
	ERR_CREATE_ROLE                 = errors.New("failed to create new role to database")
	ERR_DELETE_CREATOR              = errors.New("can't delete this creator")
	ERR_CREATE_CREATOR_DATABASE     = errors.New("failed to register new creator to database")
	ERR_DELETE                      = errors.New("failed to delete data from database")
	ERR_GET_BAD_REQUEST_ID          = errors.New("can't find any data with this id")
	ERR_CREATE_ARTICLE_DATABASE     = errors.New("failed to create new article")
	ERR_CREATE_CATEGORY_DATABASE    = errors.New("failed to create new category")
	ERR_CREATE_PRODUCT_DATABASE     = errors.New("failed to create new product")
	ERR_GET_SHIPPING_BAD_REQUEST_ID = errors.New("can't find any shipping method with this id")
	ERR_DELETE_SHIPPING             = errors.New("can't delete this shipping method")
)
