package message

import "sync"

var receivedCount int = 0
var receivedLock sync.RWMutex
var sentCount int = 0
var sentLock sync.RWMutex

// AddReceivedCount increases the received count by one
func AddReceivedCount() {
	receivedLock.Lock()
	receivedCount++
	receivedLock.Unlock()
}

// GetReceivedCount returns the current received count
func GetReceivedCount() (result int) {
	receivedLock.RLock()
	result = receivedCount
	receivedLock.RUnlock()
	return result
}

// AddSentCount increases the sent count by one
func AddSentCount() {
	sentLock.Lock()
	sentCount++
	sentLock.Unlock()
}

// GetSentCount returns the current sent count
func GetSentCount() (result int) {
	sentLock.RLock()
	result = sentCount
	sentLock.RUnlock()
	return result
}
