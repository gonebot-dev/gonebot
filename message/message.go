package message

import (
	"encoding/json"
	"log"
	"reflect"
)

// This describes a simple part of a message
type MessageSegment struct {
	// Message type
	Type string `json:"type"`
	// Which adapter is this message type for?
	// Leave it empty if you are using universal message type
	// If you have multiple adapters that have the same name, what can i say?
	Adapter string `json:"adapter"`
	// Use a message serializer to decode this
	Data string `json:"data"`
}

// Implement this to create a message serializer
type MessageSerializer interface {
	// Which adapter is this serializer for
	AdapterName() string
	// Which message type is this serializer for
	TypeName() string
	// Serialize a data to string
	Serialize(data any, messageType reflect.Type) string
	// Deserialize a message from string, will change to serializer type
	Deserialize(str string, messageType reflect.Type) any
	// Convert this message segment to raw text
	ToRawText(msg MessageSegment) string
}

type MessageType struct{}

// Serialize a data to string
func (serializer MessageType) Serialize(data any, messageType reflect.Type) string {
	var mapData map[string]any
	if reflect.TypeOf(data) != reflect.TypeOf(make(map[string]any)) {
		bs, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("[GONEBOT] | %s.Serialize: Cannot convert data to map[string]any %#v\n", messageType.String(), data)
		}
		err = json.Unmarshal(bs, &mapData)
		if err != nil {
			log.Fatalf("[GONEBOT] | %s.Serialize: Cannot convert data to map[string]any %#v\n", messageType.String(), data)
		}
	} else {
		mapData = data.(map[string]any)
	}
	value := reflect.New(messageType).Elem()
	for i := 0; i < messageType.NumField(); i++ {
		fieldName := messageType.Field(i).Tag.Get("json")
		val, ok := mapData[fieldName]
		if !ok {
			continue
		}
		fieldValue := value.Field(i)
		if !fieldValue.CanSet() {
			continue
		}
		fieldValue.Set(reflect.ValueOf(val))
	}
	result, _ := json.Marshal(value.Interface())
	return string(result)
}

// Deserialize a message from string, will change to serializer type
func (serializer MessageType) Deserialize(data string, messageType reflect.Type) any {
	if data == "" {
		data = "{}"
	}
	value := reflect.New(messageType).Elem()
	err := json.Unmarshal([]byte(data), value.Addr().Interface())
	if err != nil {
		log.Fatalf("[GONEBOT] | %s.Deserialize: Invalid data %s\n", messageType.String(), data)
	}
	return value.Interface()
}

// This describes the whole message
type Message struct {
	// Is the message to me?
	IsToMe bool
	// Which group is it in?(Only useful with group message and notice)
	Group string
	// Who is sending this message?
	Sender string
	// Whi is receiving this message?
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
func (m *Message) AttachSegment(seg MessageSegment, serializer MessageSerializer) {
	if seg.Type != serializer.TypeName() || seg.Adapter != serializer.AdapterName() {
		log.Fatalf("[GONEBOT] | Message: Invalid serializer for segment %#v\n", seg)
	}
	m.segments = append(m.segments, seg)
	m.rawText += serializer.ToRawText(seg)
}

// Text attachs a plain text message segment to message
func (m *Message) Text(text string) *Message {
	serializer := GetSerializer("text", "")
	m.AttachSegment(MessageSegment{
		Type: "text",
		Data: serializer.Serialize(TextType{
			Text: "Hello, world!",
		}, reflect.TypeOf(serializer)),
	}, serializer)
	return m
}

// Image attachs an image message segment to message
func (m *Message) Image(file string) *Message {
	serializer := GetSerializer("image", "")
	m.AttachSegment(MessageSegment{
		Type: "image",
		Data: serializer.Serialize(ImageType{
			File: file,
		}, reflect.TypeOf(serializer)),
	}, serializer)
	return m
}

// Voice attachs a voice message segment to message
func (m *Message) Voice(file string) *Message {
	serializer := GetSerializer("voice", "")
	m.AttachSegment(MessageSegment{
		Type: "voice",
		Data: serializer.Serialize(VoiceType{
			File: file,
		}, reflect.TypeOf(serializer)),
	}, serializer)
	return m
}

// Video attachs a video message segment to message
func (m *Message) Video(file string) *Message {
	serializer := GetSerializer("video", "")
	m.AttachSegment(MessageSegment{
		Type: "video",
		Data: serializer.Serialize(VideoType{
			File: file,
		}, reflect.TypeOf(serializer)),
	}, serializer)
	return m
}

// File attachs a file message segment to message
func (m *Message) File(file string) *Message {
	serializer := GetSerializer("file", "")
	m.AttachSegment(MessageSegment{
		Type: "file",
		Data: serializer.Serialize(FileType{
			File: file,
		}, reflect.TypeOf(serializer)),
	}, serializer)
	return m
}

// Attach the message contents together
func (m *Message) Join(msg Message) *Message {
	m.segments = append(m.segments, msg.segments...)
	m.rawText += msg.rawText
	return m
}

// AnySegment attachs any message segment to message
func (m *Message) AnySegment(data MessageSerializer) *Message {
	serializer := GetSerializer(data.TypeName(), data.AdapterName())
	m.AttachSegment(MessageSegment{
		Type: data.TypeName(),
		Data: serializer.Serialize(data, reflect.TypeOf(serializer)),
	}, serializer)
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

func Init() {
	// Register the builtin universal serializers
	RegisterSerializer(TextType{})
	RegisterSerializer(ImageType{})
	RegisterSerializer(VoiceType{})
	RegisterSerializer(VideoType{})
	RegisterSerializer(FileType{})
}
