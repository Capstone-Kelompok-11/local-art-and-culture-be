package response

type Files struct {
	Id        uint   `json:"id"`
	SourceId  uint   `json:"source_id"`
	SourceStr string `json:"source_str"`
	Image  	  string `json:"image"`
}
