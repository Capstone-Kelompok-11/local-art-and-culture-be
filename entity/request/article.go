package request

type Article struct {
	Id      uint       `json:"id"`
	Title   string     `json:"title"`
	Content string     `json:"content"`
	AdminId uint       `json:"adminId"`
	FilesId *uint      `json:"filesId"`
	Status  string     `json:"status"`
	Admin   SuperAdmin `json:"admin"`
}
