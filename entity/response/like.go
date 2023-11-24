package response

type Like struct {
	id        int  `json:"id"`
	UserId    int  `json:"like"`
	ArticleId uint `json:"articleId"`
	Like      bool
}
