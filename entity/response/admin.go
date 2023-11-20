package response

type Admin struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	PhoneNumber string `json:"phoneNumber"`
}
