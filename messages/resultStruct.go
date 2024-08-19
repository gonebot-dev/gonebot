package messages

import (
	"container/list"
	"sync"
)

type ResultStruct struct {
	MessageType string //"private" or "group". If not set, simply reply
	To          string //Send result to target uid. If not set, simply reply
	Text        string //Plain text message.
}

type ResultQueueStruct struct {
	mutex      sync.Mutex
	queue      *list.List
	bufferSize int
}
