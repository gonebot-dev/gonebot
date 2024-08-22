package configuations

import "os"

var Nickname string

func init() {
	Nickname = os.Getenv("NICKNAME")
	if Nickname == "" {
		Nickname = "bot"
	}
}
