package response

import "time"

type User struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"firstName"`
	LastName  	string 		`json:"lastName"`
	Email     	string 		`json:"email"`
	Token     	string 		`json:"token"`
	// AddressID   *uint     `json:"addressID"`
	PhoneNumber string    	`json:"phoneNumber"`
	NIK			string 		`json:"nik"`
	Gender		string		`json:"gender"`
	BirthDate   time.Time 	`json:"birthday"`
}

type UserCreatorResponse struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"firstName"`
	LastName  	string 		`json:"lastName"`
	Email     	string 		`json:"email"`
	Token     	string 		`json:"token"`
	// AddressID   *uint     `json:"addressID"`
	PhoneNumber string    	`json:"phoneNumber"`
	NIK			string 		`json:"nik"`
	Gender		string		`json:"gender"`
	BirthDate   time.Time 	`json:"birthday"`
	Creator     Creator
}