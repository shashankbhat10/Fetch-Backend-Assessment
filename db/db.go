package db

import (
	"github.com/google/uuid"
	"github.com/shashankbhat10/Fetch-Backend-Assessment/models"
)

type DB struct {
	Receipts map[uuid.UUID]models.Receipt
}

func Init() DB {
	store := DB{}
	store.Receipts = make(map[uuid.UUID]models.Receipt)

	return store
}
