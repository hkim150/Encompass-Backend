package util

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const (
	projectDirName = "Encompass"
)

func LoadEnv(envFile string) {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/` + envFile)
	if err != nil {
		log.Fatal("Failed to load env file: ", err)
		os.Exit(-1)
	}
}
