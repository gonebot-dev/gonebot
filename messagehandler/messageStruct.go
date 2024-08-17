package messagehandler

type messageStruct struct {
	MessageType string // "group" or "private"
	Text        string
	IsToMe      bool
	Imgs        []string
}
