package response

type Files struct {
	Id        uint   `json:"id"`
	SourceId  uint   `json:"sourceId"`
	SourceStr string `json:"sourceStr"`
	Image  	  string `json:"image"`
}
