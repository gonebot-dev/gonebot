package message

import (
	"encoding/json"
)

// This describes a simple part of a message
type MessageSegment struct {
	// Message type
	Type string `json:"type"`
	// Make sure it implements MessageType interface
	Data MessageType `json:"data"`
}

// Implement this to create a message type
type MessageType interface {
	// Which adapter is this message for
	AdapterName() string
	// Which message type is this message for
	TypeName() string
	// Convert this message segment to raw text
	ToRawText(msg MessageSegment) string
}

// This describes the whole message
type Message struct {
	// Is the message to me?
	IsToMe bool
	// Which group is it in?(Only useful with group message and notice)
	Group string
	// Who is sending this message?
	Sender string
	// Who is receiving this message?
	Receiver string
	// Who am i?
	Self string
	// All the message segments
	segments []MessageSegment
	// All the text segments are added together
	rawText string
}

func (m Message) GetSegments() []MessageSegment {
	return m.segments
}

func (m Message) GetRawText() string {
	return m.rawText
}

func NewMessage() *Message {
	return &Message{
		IsToMe:   false,
		segments: make([]MessageSegment, 0),
		rawText:  "",
	}
}

func NewReply(m Message) *Message {
	return &Message{
		IsToMe:   false,
		Group:    m.Group,
		Sender:   m.Self,
		Receiver: m.Sender,
		Self:     m.Self,
		segments: make([]MessageSegment, 0),
		rawText:  "",
	}
}

// Attach a segment for a message
func (m *Message) AttachSegment(seg MessageSegment) {
	m.segments = append(m.segments, seg)
	m.rawText += seg.Data.ToRawText(seg)
}

// Text attachs a plain text message segment to message
func (m *Message) Text(text string) *Message {
	m.AttachSegment(MessageSegment{
		Type: "text",
		Data: TextType{
			Text: text,
		},
	})
	return m
}

// Image attachs an image message segment to message
func (m *Message) Image(file string) *Message {
	m.AttachSegment(MessageSegment{
		Type: "image",
		Data: ImageType{
			File: file,
		},
	})
	return m
}

// Voice attachs a voice message segment to message
func (m *Message) Voice(file string) *Message {
	m.AttachSegment(MessageSegment{
		Type: "voice",
		Data: VoiceType{
			File: file,
		},
	})
	return m
}

// Video attachs a video message segment to message
func (m *Message) Video(file string) *Message {
	m.AttachSegment(MessageSegment{
		Type: "video",
		Data: VideoType{
			File: file,
		},
	})
	return m
}

// File attachs a file message segment to message
func (m *Message) File(file string) *Message {
	m.AttachSegment(MessageSegment{
		Type: "file",
		Data: FileType{
			File: file,
		},
	})
	return m
}

// Attach the message contents together
func (m *Message) Join(msg Message) *Message {
	m.segments = append(m.segments, msg.segments...)
	m.rawText += msg.rawText
	return m
}

// AnySegment attachs any message segment to message
func (m *Message) Any(data MessageType) *Message {
	m.AttachSegment(MessageSegment{
		Type: data.TypeName(),
		Data: data,
	})
	return m
}

func (m *Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		IsToMe   bool             `json:"is_to_me"`
		Group    string           `json:"group"`
		Sender   string           `json:"sender"`
		Receiver string           `json:"receiver"`
		Self     string           `json:"self"`
		Segments []MessageSegment `json:"segments"`
		RawText  string           `json:"raw_text"`
	}{
		IsToMe:   m.IsToMe,
		Group:    m.Group,
		Sender:   m.Sender,
		Receiver: m.Receiver,
		Self:     m.Self,
		Segments: m.segments,
		RawText:  m.rawText,
	})
}

func (m *Message) UnmarshalJSON(data []byte) error {
	tmp := &struct {
		IsToMe   bool             `json:"is_to_me"`
		Group    string           `json:"group"`
		Sender   string           `json:"sender"`
		Receiver string           `json:"receiver"`
		Self     string           `json:"self"`
		Segments []MessageSegment `json:"segments"`
		RawText  string           `json:"raw_text"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	m.IsToMe = tmp.IsToMe
	m.Group = tmp.Group
	m.Sender = tmp.Sender
	m.Receiver = tmp.Receiver
	m.Self = tmp.Self
	m.segments = tmp.Segments
	m.rawText = tmp.RawText
	return nil
}

const MESSAGE_CHAN_CAPACITY = 64
const LOG_MESSAGE_LEN_THRESHOLD = 256

type MessageChannel struct {
	channel chan Message
}

// Create a new MessageChannel
func NewMessageChannel() *MessageChannel {
	return &MessageChannel{
		channel: make(chan Message, MESSAGE_CHAN_CAPACITY),
	}
}

// Push a message to the channel
//
// # If the channel is full, the oldest message will be dropped
//
// If isReceive is true, the message will be counted as received, otherwise it will be counted as result
func (mc *MessageChannel) Push(msg Message, isReceive bool) {
	// channel full, drop
	if cap((*mc).channel) == len((*mc).channel) {
		<-(*mc).channel
	}
	// push
	(*mc).channel <- msg
	if isReceive {
		AddReceivedCount()
	} else {
		AddSentCount()
	}
}

// Pull a message from the channel
func (mc *MessageChannel) Pull() Message {
	return <-(*mc).channel
}
