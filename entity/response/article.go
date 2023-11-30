package response

import "time"

type Article struct {
	Id        uint       `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	AdminId   uint       `json:"adminId"`
	Admin     SuperAdmin `json:"admin"`
	TotalLike uint       `json:"totalLike"`
	CreatedAt time.Time  `json:"postedAt"`
	FilesId   *uint   
	Status    string     `json:"status"`
	Like      []Like     `json:"like"`
}
