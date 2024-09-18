package configurations

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Get a configuration value from the environment(provided by .env file)
//
// The builtin dotenv settings are the following:
//
// COMMAND_START="/"
//
// NICKNAME="bot"
//
// You can override these default settings to your liking.
func GetConf(name string) (result string) {
	result = os.Getenv(name)
	return result
}

// Initialize the configurations
func Init() {
	err1 := godotenv.Load()
	myEnv, err2 := godotenv.Read()
	jsonEnv, err3 := json.MarshalIndent(myEnv, "", "  ")
	if err1 != nil || err2 != nil || err3 != nil {
		log.Fatalln("[GONEBOT] | Error loading .env file.")
	}
	log.Printf("[GONEBOT] | Configuation loaded: %s\n", jsonEnv)
}
