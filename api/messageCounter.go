package api

import "sync"

var incomingCount int = 0
var incomingLock sync.RWMutex
var resultCount int = 0
var resultLock sync.RWMutex

func AddIncomingCount() {
	incomingLock.Lock()
	incomingCount++
	incomingLock.Unlock()
}
func GetIncomingCount() (result int) {
	incomingLock.RLock()
	result = incomingCount
	incomingLock.RUnlock()
	return result
}
func AddResultCount() {
	resultLock.Lock()
	resultCount++
	resultLock.Unlock()
}
func GetResultCount() (result int) {
	resultLock.RLock()
	result = resultCount
	resultLock.RUnlock()
	return result
}
