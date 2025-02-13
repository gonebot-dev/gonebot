package message

// Gonebot Universal message struct.
type Message struct {
	// UID of sender.
	SenderID string
	// UID of the bot.
	SelfID string

	// Whether it is a group message.
	IsGroup bool
	// If the message contains @at_bot
	IsToMe bool
	// Group ID
	GroupID bool

	// Message segments
	Segments []MessageSegment
}

type MessageSegments struct {
	// Type: text,image,video,sound,action.
	Type string
	// Content of the message.
	Content string
}
