package messages

import "container/list"

var ResultQueue ResultQueueStruct

func PushResult(result ResultStruct) {
	ResultQueue.mutex.Lock()
	defer ResultQueue.mutex.Unlock()
	if ResultQueue.queue.Len() == ResultQueue.bufferSize {
		//queue full
		ResultQueue.queue.Remove(ResultQueue.queue.Front())
	}
	ResultQueue.queue.PushBack(result)
}
func PopResult() (ResultStruct, bool) {
	ResultQueue.mutex.Lock()
	defer ResultQueue.mutex.Unlock()
	if ResultQueue.queue.Len() > 0 {
		result, _ := ResultQueue.queue.Front().Value.(ResultStruct)
		ResultQueue.queue.Remove(ResultQueue.queue.Front())
		return result, true
	}
	return ResultStruct{}, false
}

func init() {
	ResultQueue.queue = list.New()
	ResultQueue.bufferSize = 32
}
