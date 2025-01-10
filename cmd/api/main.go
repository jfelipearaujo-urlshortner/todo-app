package main

import (
	"fmt"

	"github.com/jfelipearaujo-urlshortner/todo-app/internal/external/http/server"
)

func main() {
	server := server.New()
	defer server.App.Shutdown()

	fmt.Println("Server is running on port :3000")

	if err := server.App.Listen(":3000"); err != nil {
		panic(err)
	}
}
