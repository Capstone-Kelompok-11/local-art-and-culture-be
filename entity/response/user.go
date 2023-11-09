package response

import "time"

type UserResponse struct {
	ID		  uint		`json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	AlamatID  *uint     `json:"alamatID"`
	NoHP      string    `json:"no_hp"`
	BirthDate time.Time `json:"birthday"`
}