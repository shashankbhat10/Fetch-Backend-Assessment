package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"github.com/shashankbhat10/Fetch-Backend-Assessment/db"
	"github.com/shashankbhat10/Fetch-Backend-Assessment/models"
)

type Server struct {
	app   *fiber.App
	store map[uuid.UUID]models.Receipt
	port  string
}

func NewServer() (*Server, error) {
	server := &Server{
		store: db.Init().Receipts,
		// Can be read from environment variable
		port: "8082",
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST",
	}))
	app.Use(logger.New(logger.Config{Format: "${status} - ${method} ${path}\n"}))

	// Routes
	app.Route("/receipts", server.ReceiptRoutes)

	server.app = app

	return server, nil
}

func (server *Server) Start() error {
	return server.app.Listen(":" + server.port)
}
