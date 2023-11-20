package response

type SuperAdmin struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	PhoneNumber string `json:"phoneNumber"`
	//Articles []Article
}

type SuperAdminArticle struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	PhoneNumber string `json:"phoneNumber"`
}