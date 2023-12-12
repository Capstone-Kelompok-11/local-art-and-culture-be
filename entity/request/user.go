package request

import "time"

type User struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"first_name" form:"first_name"`
	LastName  	string 		`json:"last_name" form:"lastName"`
	Username	string		`json:"username" form:"username"`
	Email     	string 		`json:"email" form:"email"`
	Password  	string 		`json:"password" form:"password"`
	// AddressID   *uint     `json:"addressID" form:"addressID"`
	PhoneNumber string    	`json:"phone_number" form:"phone_number"`
	NIK			string 		`json:"nik" form:"nik"`
	Gender		string		`json:"gender" form:"gender"`
	BirthDate   time.Time 	`json:"birthday" form:"birthday"`
	RoleId 		uint		`json:"role_id"`
	Role Role
}