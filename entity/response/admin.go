package response

type SuperAdmin struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	PhoneNumber string `json:"phoneNumber"`
	Role		string `json:"role"`
}