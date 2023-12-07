package response

import "time"

type User struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"firstName"`
	LastName  	string 		`json:"lastName"`
	Username	string		`json:"username"`
	Email     	string 		`json:"email"`
	Token     	string 		`json:"token"`
	// AddressID   *uint     `json:"addressID"`
	PhoneNumber string    	`json:"phoneNumber"`
	NIK			string 		`json:"nik"`
	Gender		string		`json:"gender"`
	BirthDate   time.Time 	`json:"birthday"`
	RoleId 		uint		`json:"roleId"`
	Role Role				`json:"role"`
}

type UserCreatorResponse struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"firstName"`
	LastName  	string 		`json:"lastName"`
	Username	string		`json:"username"`
	Email     	string 		`json:"email"`
	Token     	string 		`json:"token"`
	// AddressID   *uint     `json:"addressID"`
	PhoneNumber string    	`json:"phoneNumber"`
	NIK			string 		`json:"nik"`
	Gender		string		`json:"gender"`
	BirthDate   time.Time 	`json:"birthday"`
	RoleId 		uint		`json:"roleId"`
	Role Role				`json:"role"`
	Creator     Creator		`json:"creator"`
}

type Users struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"firstName"`
	LastName  	string 		`json:"lastName"`
	Username	string		`json:"username"`
	Email     	string 		`json:"email"`
	Token     	string 		`json:"token"`
	// AddressID   *uint     `json:"addressID"`
	PhoneNumber string    	`json:"phoneNumber"`
	NIK			string 		`json:"nik"`
	Gender		string		`json:"gender"`
	BirthDate   time.Time 	`json:"birthday"`
}