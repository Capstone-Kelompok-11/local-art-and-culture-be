package request

import "time"

type UserRequest struct {
	Id				uint		`json:"id"`
	FirstName 		string		`json:"first_name" form:"first_name"`
	LastName  		string		`json:"last_name" form:"last_name"`
	Email     		string		`json:"email" form:"email"`
	Password  		string		`json:"password" form:"password"`
	AddressID  		*uint		`json:"addressID" form:"addressID"`
	PhoneNumber     string		`json:"phoneNumber" form:"phoneNumber"`
	BirthDate 		time.Time	`json:"birthday" form:"birthday"`
}