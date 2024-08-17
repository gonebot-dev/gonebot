package messagehandler

type messageStruct struct {
	messageType string // "group" or "private"
	text        string
	isToMe      bool
	imgs        []string
}
