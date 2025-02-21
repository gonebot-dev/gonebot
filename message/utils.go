package message

import "strings"

// TODO: getText() etc. utils

func (m *Message) AddTextSegment(text string) {
	m.Segments = append(m.Segments, MessageSegment{
		Type:    "text",
		Content: text,
	})
}

func (m *Message) AddImageSegment(img string) {
	m.Segments = append(m.Segments, MessageSegment{
		Type:    "image",
		Content: img,
	})
}

func (m *Message) AddVideoSegment(video string) {
	m.Segments = append(m.Segments, MessageSegment{
		Type:    "video",
		Content: video,
	})
}

func (m *Message) AddSoundSegment(sound string) {
	m.Segments = append(m.Segments, MessageSegment{
		Type:    "sound",
		Content: sound,
	})
}

func (m *Message) GetText() string {
	var builder strings.Builder
	for _, seg := range m.Segments {
		if seg.Type == "text" {
			builder.WriteString(seg.Content)
		}
	}
	return builder.String()
}
