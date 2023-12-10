package response

type SuperAdmin struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	PhoneNumber string `json:"phone_number"`
	RoleId		uint   `json:"role_id"`
	Role Role		   `json:"role"`
}