package messages

import "github.com/gonebot-dev/gonebot/api"

var ResultChannel chan ResultStruct

func PushResult(result ResultStruct) {
	//channel full
	if cap(ResultChannel) == len(ResultChannel) {
		<-ResultChannel
	}
	//push
	ResultChannel <- result
	//log.Printf("Pushing Result, %d left.\n", ResultQueue.queue.Len())
}
func PopResult() ResultStruct {
	msg := <-ResultChannel
	defer api.AddResultCount()
	return msg
}

func init() {
	ResultChannel = make(chan ResultStruct, 32)
}
