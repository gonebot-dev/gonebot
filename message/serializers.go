package message

import (
	"fmt"
	"log"
	"reflect"
)

type TextType struct {
	MessageType
	Text string `json:"text"`
}

func (serializer TextType) Matcher(typeName, adapterName string) bool {
	return typeName == "text" && adapterName == ""
}

func (serializer TextType) ToRawText(msg MessageSegment) string {
	result := serializer.Deserialize(msg.Data, reflect.TypeOf(serializer))
	return result.(TextType).Text
}

type ImageType struct {
	MessageType
	File string `json:"file"`
}

func (serializer ImageType) Matcher(typeName, adapterName string) bool {
	return typeName == "image" && adapterName == ""
}

func (serializer ImageType) ToRawText(msg MessageSegment) string {
	result := serializer.Deserialize(msg.Data, reflect.TypeOf(serializer))
	return fmt.Sprintf("[image:%s]", result.(ImageType).File)
}

type VoiceType struct {
	MessageType
	File string `json:"file"`
}

func (serializer VoiceType) Matcher(typeName, adapterName string) bool {
	return typeName == "voice" && adapterName == ""
}

func (serializer VoiceType) ToRawText(msg MessageSegment) string {
	result := serializer.Deserialize(msg.Data, reflect.TypeOf(serializer))
	return fmt.Sprintf("[voice:%s]", result.(VoiceType).File)
}

type VideoType struct {
	MessageType
	File string `json:"file"`
}

func (serializer VideoType) Matcher(typeName, adapterName string) bool {
	return typeName == "video" && adapterName == ""
}

func (serializer VideoType) ToRawText(msg MessageSegment) string {
	result := serializer.Deserialize(msg.Data, reflect.TypeOf(serializer))
	return fmt.Sprintf("[video:%s]", result.(VideoType).File)
}

type FileType struct {
	MessageType
	File string `json:"file"`
}

func (serializer FileType) Matcher(typeName, adapterName string) bool {
	return typeName == "file" && adapterName == ""
}

func (serializer FileType) ToRawText(msg MessageSegment) string {
	result := serializer.Deserialize(msg.Data, reflect.TypeOf(serializer))
	return fmt.Sprintf("[file:%s]", result.(FileType).File)
}

var SerializerRegistry = make(map[string]MessageSerializer)

func RegisterSerializer(typeName, adapterName string, serializer MessageSerializer) {
	identifier := fmt.Sprintf("%s:%s", adapterName, typeName)
	if _, exists := SerializerRegistry[identifier]; exists {
		log.Fatalf("[GONEBOT] | RegisterSerializer: Duplicate serializer for type %s", identifier)
	}
	SerializerRegistry[identifier] = serializer
}

func GetSerializer(typeName, adapterName string) MessageSerializer {
	matcher := fmt.Sprintf("%s:%s", adapterName, typeName)
	for key, serializer := range SerializerRegistry {
		if key == matcher {
			return serializer
		}
	}
	log.Printf("[GONEBOT] | GetSerializer: No serializer for identifier %s", matcher)
	return nil
}
