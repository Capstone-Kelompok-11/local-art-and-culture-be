package request

type Files struct {
	Id        uint   `json:"id" form:"id"`
	SourceId  uint   `json:"source_id" form:"source_id"`
	SourceStr string `json:"source_str" form:"source_str"`
	Image  	  string `json:"image" form:"image"`
}
