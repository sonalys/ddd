package application

import (
	"context"

	"github.com/sonalys/ddd/domain/entity"
	"github.com/sonalys/ddd/domain/repository"
	"github.com/sonalys/ddd/infraestructure/persistence"
)

// cartApp implements the Cart application layer.
type cartApp struct {
	repository.Cart
}

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

func NewCartApp(ctx context.Context, m *persistence.Mongo) (Cart, error) {
	persistence, err := persistence.NewCartPersistence(ctx, m)
	if err != nil {
		return nil, err
	}

	return &cartApp{
		Cart: persistence,
	}, nil
}

func (c cartApp) AddItem(ctx context.Context, cartID string, item *entity.CartItem) error {
	if err := item.Valid(); err != nil {
		return err
	}

	c.Cart.AddItem(ctx, cartID, item)
	return nil
}
