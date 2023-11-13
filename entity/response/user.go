package response

import "time"

type UserResponse struct {
	Id		  		uint	  `json:"id"`
	FirstName 		string    `json:"firstName"`
	LastName  		string    `json:"lastName"`
	Email     		string    `json:"email"`
	Token       	string 	  `json:"token"`
	AddressID  		*uint     `json:"addressID"`
	PhoneNumber      string   `json:"phoneNumber"`
	BirthDate 		time.Time `json:"birthday"`
}
