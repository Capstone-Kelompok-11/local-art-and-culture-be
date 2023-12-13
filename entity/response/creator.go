package response

type Creators struct {
	Id          uint   `json:"id"`
	OutletName  string `json:"outlet_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserId      uint   `json:"user_id"`
	Roles     	string   `json:"roles"`
	AddressId   *uint  `json:"address_id"`
	Users       Users   `json:"users"`
	Role       Role   `json:"role"`
}

type Creator struct {
	Id          uint   `json:"id"`
	OutletName  string `json:"outlet_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserId      uint   `json:"user_nd"`
	Roles     	string  `json:"roles"`
	AddressId   *uint  `json:"address_id"`
}