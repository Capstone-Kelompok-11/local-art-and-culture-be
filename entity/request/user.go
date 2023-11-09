package request

import "time"

type UserRequest struct {
	FirstName string	`json:"first_name" form:"first_name"`
	LastName  string	`json:"last_name" form:"last_name"`
	Email     string	`json:"email" form:"email"`
	Password  string	`json:"password" form:"password"`
	AlamatID  *uint		`json:"alamatID" form:"alamatID"`
	NoHP      string	`json:"no_hp" form:"no_hp"`
	BirthDate time.Time	`json:"birthday" form:"birthday"`
}