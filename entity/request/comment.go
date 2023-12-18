package request

type Comment struct {
	Id			uint	`json:"id"`
	Comment		string	`json:"comment"`
	ArticleId	uint	`json:"article_id"`
	LikeId		*uint	`json:"like_id"`
	UserId		uint	`json:"user_id"`		
}