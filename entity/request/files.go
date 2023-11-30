package request

type Files struct {
	Id        uint   `json:"id" form:"id"`
	SourceId  uint   `json:"sourceId" form:"sourceId"`
	SourceStr string `json:"sourceStr" form:"sourceStr"`
	Image  	  string `json:"image" form:"image"`
}
