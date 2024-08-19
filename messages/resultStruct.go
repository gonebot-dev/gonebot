package messages

type ResultStruct struct {
	To   string //Send result to target uid. If not set, reply the sender.
	Text string //Plain text message.
}
