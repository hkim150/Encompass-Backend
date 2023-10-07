package main

import (
	"encompass/api"
	"encompass/util"
	"os"
)

func main() {
	util.LoadEnv(".env")

	server := api.NewServer()
	server.Logger.Fatal(server.Start(os.Getenv("DB_DRIVER")))
}
