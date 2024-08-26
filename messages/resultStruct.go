package messages

type ResultStruct struct {
	MessageType string   // "private" or "group". If not set, simply reply
	To          string   // Send result to target uid. If not set, simply reply
	Text        string   // Plain text message.
	Imgs        []string //!Not implemented
}
