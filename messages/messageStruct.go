package messages

import (
	"container/list"
	"sync"
)

// An unified message struct in gonebot.
type MessageStruct struct {
	MessageType string // "group" or "private"
	Text        string
	IsToMe      bool
	Imgs        []string
}

type MessageQueueStruct struct {
	queue      *list.List
	mutex      sync.Mutex
	bufferSize int
}
