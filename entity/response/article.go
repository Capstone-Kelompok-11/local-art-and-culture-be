package response

type Article struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	AdminId uint   `json:"adminId"`
}
