package request

type Creator struct {
	Id          uint   	`json:"id"`
	OutletName  string 	`json:"outlet_name"`
	Email       string 	`json:"email"`
	PhoneNumber string 	`json:"phone_number"`
	Roles    	string  `json:"roles"`
	UserId 		uint
	RoleId 		uint
	AddressId   *uint  	`json:"address_id"`
	Address 	string	`json:"address"`
	Role       	Role   	`json:"role"`
}