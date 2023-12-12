package response

import "time"

type User struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"first_name"`
	LastName  	string 		`json:"last_name"`
	Username	string		`json:"username"`
	Email     	string 		`json:"email"`
	Token     	string 		`json:"token"`
	// AddressID   *uint     `json:"addressID"`
	PhoneNumber string    	`json:"phone_number"`
	NIK			string 		`json:"nik"`
	Gender		string		`json:"gender"`
	BirthDate   time.Time 	`json:"birthday"`
	Status 		string		`json:"status"`
	DeletedAt   time.Time	`json:"deleted_at,omitempty"`
	Date		string	    `json:"date"`
	CreatedAt   time.Time	`json:"created_at,omitempty"`
	RoleId 		uint		`json:"role_id"`
	Role Role				`json:"role"`
}

type UserCreatorResponse struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"first_name"`
	LastName  	string 		`json:"last_name"`
	Username	string		`json:"username"`
	Email     	string 		`json:"email"`
	Token     	string 		`json:"token"`
	// AddressID   *uint     `json:"addressID"`
	PhoneNumber string    	`json:"phone_number"`
	NIK			string 		`json:"nik"`
	Gender		string		`json:"gender"`
	BirthDate   time.Time 	`json:"birthday"`
	RoleId 		uint		`json:"role_id"`
	Role Role				`json:"role"`
	Creator     Creator		`json:"creator"`
}

type Users struct {
	Id        	uint   		`json:"id"`
	FirstName 	string 		`json:"first_name"`
	LastName  	string 		`json:"last_name"`
	Username	string		`json:"username"`
	Email     	string 		`json:"email"`
	Token     	string 		`json:"token"`
	// AddressID   *uint     `json:"addressID"`
	PhoneNumber string    	`json:"phone_number"`
	NIK			string 		`json:"nik"`
	Gender		string		`json:"gender"`
	BirthDate   time.Time 	`json:"birthday"`
}