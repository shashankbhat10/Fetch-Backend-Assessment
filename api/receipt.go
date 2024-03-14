package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/shashankbhat10/Fetch-Backend-Assessment/models"
	"github.com/shashankbhat10/Fetch-Backend-Assessment/util"
)

func (server *Server) ReceiptRoutes(router fiber.Router) {
	// Add receipt to DB and return receiptID
	router.Post("/process", server.StoreReceipt)
	// Get total points earned by a receipt
	router.Get("/:id/points", server.GetPoints)
}

func (server *Server) StoreReceipt(ctx *fiber.Ctx) error {
	body := ctx.Body()
	if body == nil {
		log.Println("Error while accessing request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "error while reading request body"})
	}

	req := models.Receipt{}
	err := json.Unmarshal(body, &req)
	if err != nil {
		log.Println("Error while accessing request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "error while parsing request body"})
	}

	newId := uuid.New()
	server.store[newId] = req

	return ctx.Status(fiber.StatusOK).JSON(models.PostReceiptResponse{Id: newId.String()})
}

func (server *Server) GetPoints(ctx *fiber.Ctx) error {
	receiptId := ctx.Params("id")
	if receiptId == "" {
		log.Println("Receipt ID is empty")
		return ctx.Status(http.StatusNotFound).JSON(models.ErrorResponse{Error: "Receipt ID is empty"})
	}

	parsedReceiptId, err := uuid.Parse(receiptId)
	if err != nil {
		log.Println("Invalid Receipt ID", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Invalid Receipt ID"})
	}

	receipt, keyPresent := server.store[parsedReceiptId]
	if !keyPresent {
		log.Println("No Receipt found for ReceiptID")
		return ctx.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Error: "No receipt found for given receipt ID"})
	}

	points, err := util.GetPoints(receipt)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.GetPointsResponse{Points: points})
}
