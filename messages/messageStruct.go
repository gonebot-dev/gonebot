package messages

// An unified message struct in gonebot.
type MessageStruct struct {
	MessageType string // "group" or "private"
	Text        string
	IsToMe      bool
	Imgs        []string
}
