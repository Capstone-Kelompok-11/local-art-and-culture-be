package request

type SuperAdmin struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	RoleId      uint   `json:"role_id"`
	Role        Role
}
