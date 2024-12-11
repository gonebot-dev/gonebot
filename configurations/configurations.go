package configurations

import (
	"encoding/json"

	"github.com/gonebot-dev/gonebot/logging"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

// Initialize the configurations
func Init() {
	err1 := godotenv.Load()
	myEnv, err2 := godotenv.Read()
	jsonEnv, err3 := json.MarshalIndent(myEnv, "", "  ")
	if err1 != nil || err2 != nil || err3 != nil {
		logging.Log(zerolog.FatalLevel, "GoneBot", "Error loading .env file.")
	}
	logging.Logf(zerolog.InfoLevel, "GoneBot", "Configuation loaded: %s\n", jsonEnv)
}
