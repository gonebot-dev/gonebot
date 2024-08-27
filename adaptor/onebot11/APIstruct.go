package onebot11

type APIPayload struct {
	Action string `json:"action"`
	Params any    `json:"params"`
}

type APISendPrivateMessage struct {
	UserID  int           `json:"user_id"`
	Message []interface{} `json:"message"` //MessageSegment*s
}
type APISendGroupMessage struct {
	GroupID int           `json:"group_id"`
	Message []interface{} `json:"message"` //MessageSegment*s
}

type MessageSegmentText struct {
	Type string `json:"type"` //must be text
	Data struct {
		Text string `json:"text"` //text content
	} `json:"data"`
}
type MessageSegmentImg struct {
	Type string `json:"type"`
	Data struct {
		Uri string `json:"file"`
	} `json:"data"`
}
