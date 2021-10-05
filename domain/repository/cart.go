package repository

import (
	"context"

	"github.com/sonalys/ddd/domain/entity"
)

// Cart is a repository to manage cart entities.
type Cart interface {
	// AddItem adds a new item for an existent cart.
	AddItem(ctx context.Context, cartID string, item *entity.CartItem) error
	// Create instantiates a new cart in the persistence, and returns the entity.
	Create(ctx context.Context) (*entity.Cart, error)
	// Get returns a cart with all it's items.
	Get(ctx context.Context, cartID string) (*entity.Cart, error)
	// Remove removes a cart from the persistence.
	Remove(ctx context.Context, cartID string) error
	// RemoveItem removes an item from an existent cart.
	RemoveItem(ctx context.Context, cartID, itemID string) error
}
