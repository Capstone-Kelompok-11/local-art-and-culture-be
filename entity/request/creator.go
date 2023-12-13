package request

type Creator struct {
	Id          uint   `json:"id"`
	OutletName  string `json:"outlet_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserId      uint   
	Roles    	string   `json:"roles"`
	AddressId   *uint  `json:"address_id"`
	Users       User   `json:"users"`
	Role       Role   `json:"role"`
}
