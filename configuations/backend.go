package configuations

import "os"

var BackendHostAddress string //default "127.0.0.1:21390"

func init() {
	BackendHostAddress = os.Getenv("HOST")
	if BackendHostAddress == "" {
		BackendHostAddress = "127.0.0.1:21390"
	}
}
