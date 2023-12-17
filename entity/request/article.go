package request

type Article struct {
	Id      uint       `json:"id"`
	Title   string     `json:"title"`
	Content string     `json:"content"`
	AdminId uint       `json:"admin_id"`
	FilesId *uint      `json:"files_id"`
	Status  string     `json:"status"`
	Admin   SuperAdmin `json:"admin"`
}
