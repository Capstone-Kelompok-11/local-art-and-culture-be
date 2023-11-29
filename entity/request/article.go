package request

type Article struct {
	Id      	uint   `json:"id"`
	Title   	string `json:"title"`
	AdminId 	uint   `json:"adminId"`
	Content 	string `json:"content"`
	PictureId 	uint   `json:"pictureId"`
	Status		string `json:"status"`
}
