package request

type Article struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	AdminId uint   `json:"adminId"`
	Admin   Admin  `json"admin"`
}
