package response

type Like struct {
	Id        uint `json:"id"`
	UserId    uint `json:"userId"`
	ArticleId uint `json:"articleId"`
	Active    bool `json:"like"`
}
