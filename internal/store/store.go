package store

import (
	"context"

	"github.com/MalukiMuthusi/cryptocurrencies/internal/models"
)

// Store encapsulates the storage used by the service
type Store interface {
	List(c context.Context) ([]*models.ListRes, error)
}
