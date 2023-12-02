package request

type Comment struct {
	Id			uint	`json:"id"`
	Comment		string	`json:"comment"`
	ArticleId	uint	`json:"articleId"`
	LikeId		*uint	`json:"likeId"`
	UserId		uint	`json:"userId"`		
}