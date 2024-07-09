package settings

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


var(
	// DbConnectionString is the connection string with database
	DbConnectionString = ""
	// DbName is the name of the database
	DbName = ""
	// Port is the API port 
	Port = 0

)

// LoadSettings will load all environments
func LoadSettings(){
	var erro error

	if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	DbConnectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s", 
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_SERVER"),
		os.Getenv("DB_PORT"),
	)

	DbName = os.Getenv("DB_NAME")


}