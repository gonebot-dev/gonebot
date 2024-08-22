package configuations

import "os"

var GlobalPrefix string

func init() {
	GlobalPrefix = os.Getenv("COMMAND_START")
	if GlobalPrefix == "" {
		GlobalPrefix = "/"
	}
}
