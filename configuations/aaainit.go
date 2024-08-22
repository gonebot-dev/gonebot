package configuations

import (
	"encoding/json"
	"log"

	"github.com/joho/godotenv"
)

// Why named as aconfiguation? Multiple init() func will be invoked in dictionary order!
func init() {
	err1 := godotenv.Load()
	myEnv, err2 := godotenv.Read()
	jsonEnv, err3 := json.Marshal(myEnv)
	if err1 != nil || err2 != nil || err3 != nil {
		log.Println("Error loading .env file.")
	}
	log.Printf("Configuation loaded: %s\n", jsonEnv)
}
