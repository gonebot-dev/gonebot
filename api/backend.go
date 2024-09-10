package api

var backend = "Undefined"

func SetBackend(b string) {
	backend = b
}

func GetBackend() string {
	return backend
}
