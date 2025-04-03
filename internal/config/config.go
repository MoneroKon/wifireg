// made by recanman
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func checkEnv(name string) {
	if _, exists := os.LookupEnv(name); !exists {
		log.Fatalf("Environment variable %s not set", name)
	}
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load environment variables")
	}

	checkEnv("OPNSENSE_HOST")
	checkEnv("OPNSENSE_KEY")
	checkEnv("OPNSENSE_SECRET")
}

func OpnsenseHost() string {
	return os.Getenv("OPNSENSE_HOST")
}
func OpnsenseCert() string {
	return os.Getenv("OPNSENSE_CERT")
}
func OpnsenseKey() string {
	return os.Getenv("OPNSENSE_KEY")
}
func OpnsenseSecret() string {
	return os.Getenv("OPNSENSE_SECRET")
}

func Port() string {
	port, exists := os.LookupEnv("PORT")
	if exists {
		return port
	} else {
		return "3000"
	}
}
