package response

import "time"

type UserResponse struct {
	Id		  		uint	  `json:"id"`
	FirstName 		string    `json:"first_name"`
	LastName  		string    `json:"last_name"`
	Email     		string    `json:"email"`
	Password  		string    `json:"password"`
	AddressID  		*uint     `json:"addressID"`
	PhoneNumber      string   `json:"phoneNumber"`
	BirthDate 		time.Time `json:"birthday"`
}
