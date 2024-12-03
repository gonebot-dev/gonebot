package message

import (
	"encoding/json"
	"fmt"
)

type TextType struct {
	Text string `json:"text"`
}

func (serializer TextType) AdapterName() string {
	return ""
}

func (serializer TextType) TypeName() string {
	return "text"
}

func (serializer TextType) ToRawText(msg MessageSegment) string {
	return msg.Data.(TextType).Text
}

type ImageType struct {
	File string `json:"file"`
}

func (serializer ImageType) AdapterName() string {
	return ""
}

func (serializer ImageType) TypeName() string {
	return "image"
}

func (serializer ImageType) ToRawText(msg MessageSegment) string {
	return fmt.Sprintf("[image:%s]", msg.Data.(ImageType).File)
}

type VoiceType struct {
	File string `json:"file"`
}

func (serializer VoiceType) AdapterName() string {
	return ""
}

func (serializer VoiceType) TypeName() string {
	return "voice"
}

func (serializer VoiceType) ToRawText(msg MessageSegment) string {
	return fmt.Sprintf("[voice:%s]", msg.Data.(VoiceType).File)
}

type VideoType struct {
	File string `json:"file"`
}

func (serializer VideoType) AdapterName() string {
	return ""
}

func (serializer VideoType) TypeName() string {
	return "video"
}

func (serializer VideoType) ToRawText(msg MessageSegment) string {
	return fmt.Sprintf("[video:%s]", msg.Data.(VideoType).File)
}

type FileType struct {
	File string `json:"file"`
}

func (serializer FileType) AdapterName() string {
	return ""
}

func (serializer FileType) TypeName() string {
	return "file"
}

func (serializer FileType) ToRawText(msg MessageSegment) string {
	return fmt.Sprintf("[file:%s]", msg.Data.(FileType).File)
}

// Convert raw MessageSegment.Data to built-in MessageType
func ToBuiltIn(typeName string, msg any) MessageType {
	tmp, _ := json.Marshal(msg)
	switch typeName {
	case "text":
		var result TextType
		_ = json.Unmarshal(tmp, &result)
		return result
	case "image":
		var result ImageType
		_ = json.Unmarshal(tmp, &result)
		return result
	case "voice":
		var result VoiceType
		_ = json.Unmarshal(tmp, &result)
		return result
	case "video":
		var result VideoType
		_ = json.Unmarshal(tmp, &result)
		return result
	case "file":
		var result FileType
		_ = json.Unmarshal(tmp, &result)
		return result
	default:
		return nil
	}
}
