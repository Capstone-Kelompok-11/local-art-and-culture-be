package response

type SaveChatbot struct {
	Id			uint	`json:"id"`
	Message 	string	`json:"message"`
	Response	string	`json:"response"`
	UserId		uint	`json:"userId"`
}