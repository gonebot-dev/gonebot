package messages

// An unified message struct in gonebot.
type MessageStruct struct {
	MessageType string   // "group" or "private"
	Text        string   // Plain text content of the message.
	SenderID    string   // Uid who send the message.
	SenderName  string   // Nickname who send the message.
	IsToMe      bool     // Whether the message is sent to me.
	Imgs        []string //!Not implemented
	SelfID      string   // Uid of bot
}
