package request

type Guest struct {
	Id			uint 	`json:"id"`
	Name		string	`json:"name"`
	Role		string	`json:"role"`
	//PictureId	*uint	`json:"pictureId"`
	EventId		uint	`json:"eventId"`
}