package messages

type ResultStruct struct {
	MessageType string   // "private" or "group". If not set, simply reply
	To          string   // Target uid. If not set, simply reply
	Text        string   // Plain text message.
	Imgs        []string // Images. Currently we support absolute path 'file:///C:\\test\1.png', url 'http://...', Base64 'base64://xxx'
}
