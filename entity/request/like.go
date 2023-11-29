package request

type Like struct {
    Id        uint   `json:"id"`
    UserId    uint   `json:"userId"`
    SourceId  uint   `json:"sourceId"`
    SourceStr string `json:"sourceStr"`
    Active    bool   `json:"like"`
}
