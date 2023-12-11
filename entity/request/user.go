package request

import "time"

type User struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"firstName" form:"firstName"`
	LastName  	string 		`json:"lastName" form:"lastName"`
	Username	string		`json:"username" form:"username"`
	Email     	string 		`json:"email" form:"email"`
	Password  	string 		`json:"password" form:"password"`
	// AddressID   *uint     `json:"addressID" form:"addressID"`
	PhoneNumber string    	`json:"phoneNumber" form:"phoneNumber"`
	NIK			string 		`json:"nik" form:"nik"`
	Gender		string		`json:"gender" form:"gender"`
	BirthDate   time.Time 	`json:"birthday" form:"birthday"`
	RoleId 		uint		`json:"roleId"`
	Role Role
}

type EmailRequest struct {
	Email		string		`json:"email" form:"email"`
	OTP			string		`json:"otp" form:"otp"`
}