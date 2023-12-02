package errors

import "errors"

var (
	ERR_EMAIL_IS_EMPTY                        = errors.New("email is empty")
	ERR_MESSAGE_IS_EMPTY                      = errors.New("message is empty")
	ERR_RESPONSE_IS_EMPTY                     = errors.New("response is empty")
	ERR_IMAGE_IS_EMPTY                        = errors.New("image is empty")
	ERR_PRICE_IS_EMPTY                        = errors.New("price is empty")
	ERR_COMMENT_IS_EMPTY                      = errors.New("comment is empty")
	ERR_DESCRIPTION_IS_EMPTY                  = errors.New("description is empty")
	ERR_STATUS_IS_EMPTY                       = errors.New("status is empty")
	ERR_NAME_IS_EMPTY                         = errors.New("name is empty")
	ERR_OUTLET_NAME_IS_EMPTY                  = errors.New("outlet name is empty")
	ERR_CATEGORY_IS_EMPTY                     = errors.New("category is empty")
	ERR_TYPE_IS_EMPTY                         = errors.New("type is empty")
	ERR_PASSWORD_IS_EMPTY                     = errors.New("password is empty")
	ERR_PHONE_NUMBER_IS_EMPTY                 = errors.New("phone number is empty")
	ERR_GET_ADMIN_BAD_REQUEST_ID              = errors.New("can't find any admin with this id")
	ERR_GET_CATEGORY_BAD_REQUEST_TYPE         = errors.New("can't find any category with this type")
	ERR_GET_USER_BAD_REQUEST_ID               = errors.New("can't find any user with this id")
	ERR_DELETE_ADMIN                          = errors.New("can't delete this admin")
	ERR_DELETE_USER                           = errors.New("can't delete this user")
	ERR_REGISTER_ADMIN_DATABASE               = errors.New("failed to register new admin to database")
	ERR_BCRYPT_PASSWORD                       = errors.New("failed to bcrypt password")
	ERR_UPDATE_DATA                           = errors.New("failed to update data")
	ERR_TOKEN                                 = errors.New("failed to create new token")
	ERR_WRONG_PASSWORD                        = errors.New("wrong password")
	ERR_EMAIL_NOT_FOUND                       = errors.New("email not found")
	ERR_REGISTER_USER_DATABASE                = errors.New("failed to register new user to database")
	ERR_GET_DATA                              = errors.New("database can't request any data right now")
	ERR_ROLE_IS_EMPTY                         = errors.New("role is empty")
	ERR_GET_ROLE_BAD_REQUEST_ID               = errors.New("can't find any role with this id")
	ERR_GET_PAYMENT_BAD_REQUEST_ID            = errors.New("can't find any payment with this id")
	ERR_GET_CREATOR_BAD_REQUEST_ID            = errors.New("can't find any creator with this id")
	ERR_GET_PRODUCT_BAD_REQUEST_ID            = errors.New("can't find any product with this id")
	ERR_GET_COMMENT_BAD_REQUEST_ID            = errors.New("can't find any comment with this id")
	ERR_GET_CATEGORY_BAD_REQUEST_ID           = errors.New("can't find any category with this id")
	ERR_DELETE_ROLE                           = errors.New("can't delete this role")
	ERR_DELETE_CATEGORY                       = errors.New("can't delete this category")
	ERR_TITLE_IS_EMPTY                        = errors.New("title is empty")
	ERR_CONTENT_IS_EMPTY                      = errors.New("content is empty")
	ERR_CREATE_ROLE                           = errors.New("failed to create new role to database")
	ERR_CREATE_PAYMENT                        = errors.New("failed to create new payment to database")
	ERR_DELETE_CREATOR                        = errors.New("can't delete this creator")
	ERR_CREATE_CREATOR_DATABASE               = errors.New("failed to register new creator to database")
	ERR_DELETE                                = errors.New("failed to delete data from database")
	ERR_GET_BAD_REQUEST_ID                    = errors.New("can't find any data with this id")
	ERR_CREATE_ARTICLE_DATABASE               = errors.New("failed to create new article")
	ERR_CREATE_COMMENT_DATABASE               = errors.New("failed to create new comment")
	ERR_CREATE_CATEGORY_DATABASE              = errors.New("failed to create new category")
	ERR_CREATE_PRODUCT_DATABASE               = errors.New("failed to create new product")
	ERR_CREATE_SAVE_DATABASE                  = errors.New("failed to create new save chatbot")
	ERR_GET_SHIPPING_BAD_REQUEST_ID           = errors.New("can't find any shipping method with this id")
	ERR_GET_SAVE_BAD_REQUEST_ID               = errors.New("can't find any save chatbot method with this id")
	ERR_DELETE_SHIPPING                       = errors.New("can't delete this shipping method")
	ERR_GET_EVENT_BAD_REQUEST_ID              = errors.New("can't find any Event with this id")
	ERR_EVENT_DATE_IS_EMPTY                   = errors.New("event date is empty")
	ERR_QTY_IS_EMPTY                          = errors.New("quantity is empty")
	ERR_EVENT_SCHEDULE_IS_EMPTY               = errors.New("event schedule is empty")
	ERR_CREATE_EVENT_DATABASE                 = errors.New("failed to create new event to database")
	ERR_CREATE_TRANSACTION_DATABASE           = errors.New("failed to create new transaction to database")
	ERR_GET_TRANSACTION_BAD_REQUEST_ID        = errors.New("can't find any transaction with this id")
	ERR_DELETE_TRANSACTION                    = errors.New("can't delete this transaction")
	ERR_GET_TICKET_BAD_REQUEST_ID             = errors.New("can't find any Ticket with this id")
	ERR_CREATE_TICKET_DATABASE                = errors.New("failed to create new ticket")
	ERR_GET_ARTICLE_BAD_REQUEST_ID            = errors.New("can't find any Article with this id")
	ERR_CREATE_WISHLIST_DATABASE              = errors.New("failed to create new wishlist to database")
	ERR_CREATE_LIKE_DATABASE                  = errors.New("can't like this article")
	ERR_ULASAN_IS_EMPTY                       = errors.New("feedback is empty")
	ERR_CREATE_RATING_DATABASE                = errors.New("failed to create new feedback to database")
	ERR_CREATE_GUEST_DATABASE                 = errors.New("failed to create new guest to database")
	ERR_RATING_IS_EMPTY                       = errors.New("rating is empty")
	ERR_CREATE_TRANSACTION_DETAIL             = errors.New("failed to create new transaction detail")
	ERR_GET_TRANSACTION_DETAIL_BAD_REQUEST_ID = errors.New("can't find any Transaction Detail with this id")
	ERR_GET_FILES_BAD_REQUEST_ID              = errors.New("can't find any file with this id")
	ERR_CREATE_FILES_DATABASE                 = errors.New("failed to upload new files to database")
	ERR_SOURCE_IS_EMPTY                       = errors.New("source is empty`")
)
