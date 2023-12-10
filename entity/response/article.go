package response

import "time"

type Article struct {
	Id        uint       `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	AdminId   uint       `json:"admin_id"`
	Admin     SuperAdmin `json:"admin"`
	TotalLike uint       `json:"total_like"`
	CreatedAt time.Time  `json:"posted_at"`
	FilesId   *uint   
	Status    string     `json:"status"`
	Like      []Like     `json:"like"`
}
