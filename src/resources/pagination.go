package resources

type CursorResource struct {
	NextCursor string      `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
	Count      int         `json:"count"`
	Data       interface{} `json:"data"`
}
