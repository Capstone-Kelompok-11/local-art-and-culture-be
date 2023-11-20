package response

type Creator struct {
	Id          uint   `json:"id"`
	OutletName  string `json:"outletName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	UserId      uint   `json:"userId"`
	RoleId      uint   `json:"roleId"`
	AddressId   *uint  `json:"addressId"`
	Users       User   `json:"users"`
	Roles       Role   `json:"role"`
}
