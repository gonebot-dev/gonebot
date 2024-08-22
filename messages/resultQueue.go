package messages

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
	return msg
}

func init() {
	ResultChannel = make(chan ResultStruct, 32)
}
