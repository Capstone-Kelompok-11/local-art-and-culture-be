package request

import "time"

type User struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"firstName" form:"firstName"`
	LastName  	string 		`json:"lastName" form:"lastName"`
	Email     	string 		`json:"email" form:"email"`
	Password  	string 		`json:"password" form:"password"`
	// AddressID   *uint     `json:"addressID" form:"addressID"`
	PhoneNumber string    	`json:"phoneNumber" form:"phoneNumber"`
	NIK			string 		`json:"nik"`
	Gender		string		`json:"gender"`
	BirthDate   time.Time 	`json:"birthday" form:"birthday"`
	// Creator     Creator
}
