package response

type Creators struct {
	Id          uint   	`json:"id"`
	OutletName  string 	`json:"outlet_name"`
	Email       string 	`json:"email"`
	PhoneNumber string 	`json:"phone_number"`
	AddressId   *uint  	`json:"address_id"`
	Address 	string	`json:"address"`
	Users       Users   `json:"users"`
	Role       	Role   	`json:"role"`
}

type Creator struct {
	Id          uint   `json:"id"`
	OutletName  string `json:"outlet_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserId      uint   `json:"user_nd"`
	AddressId   *uint  `json:"address_id"`
}