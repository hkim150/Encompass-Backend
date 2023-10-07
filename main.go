package main

import (
	"encompass/api"
)

func main() {
	server := api.NewServer()
	server.Logger.Fatal(server.Start(":1323"))
}
