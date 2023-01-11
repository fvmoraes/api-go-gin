package main

import (
	"api-go-gin/initializers"
	"api-go-gin/servers"
)

func main() {
	initializers.StartDatabaseConnect()
	initializers.StartMigration()
	servers.RunMyFoobarServer()
}
