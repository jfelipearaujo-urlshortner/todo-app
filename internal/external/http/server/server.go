package server

import "github.com/gofiber/fiber/v2"

type Server struct {
	container *container

	App *fiber.App
}

func New() *Server {
	server := Server{
		container: newContainer(),
		App:       fiber.New(),
	}

	server.registerRoutes()

	return &server
}
