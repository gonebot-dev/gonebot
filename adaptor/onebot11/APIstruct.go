package onebot11

type APIPayload struct {
	Action string `json:"action"`
	Params any    `json:"params"`
}

type APISendPrivateMessage struct {
	UserID  int           `json:"user_id"`
	Message []interface{} `json:"message"`
}

type MessageSegmentText struct {
	Type string `json:"type"`
	Data struct {
		Text string `json:"text"`
	} `json:"data"`
}
