package request

type Like struct {
    Id        uint   `json:"id"`
    UserId    uint   `json:"user_id"`
    SourceId  uint   `json:"source_id"`
    SourceStr string `json:"source_str"`
    Active    bool   `json:"like"`
}
