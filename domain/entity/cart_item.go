package entity

import (
	"strings"

	"github.com/pkg/errors"
)

// CartItem represents a cart item.
type CartItem struct {
	Discount float32 `json:"discount"`
	// ID links to Item entities.
	ID       string  `json:"id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	Price    float32 `json:"price" bson:"price"`
	Quantity uint    `json:"quantity" bson:"quantity"`
}

var (
	// ErrInvalidCartItem indicates a cart item contains errors.
	ErrInvalidCartItem error = errors.New("cart item is invalid")
)

// Valid implements Validatable interface.
func (c CartItem) Valid() error {
	invalidFieldErrors := []string{}

	if c.Name == "" {
		invalidFieldErrors = append(invalidFieldErrors, "name must not be empty")
	}

	if c.Quantity == 0 {
		invalidFieldErrors = append(invalidFieldErrors, "quantity must be greater than 0")
	}

	if c.Price <= 0 {
		invalidFieldErrors = append(invalidFieldErrors, "price must be greater than 0")
	}

	if c.Discount < 0 {
		invalidFieldErrors = append(invalidFieldErrors, "discount must be greater than 0")
	}

	if c.Discount > c.Price {
		invalidFieldErrors = append(invalidFieldErrors, "discount cannot be bigger than price")
	}

	if len(invalidFieldErrors) > 0 {
		msg := strings.Join(invalidFieldErrors, ";")
		return errors.Wrap(ErrInvalidCartItem, msg)
	}
	return nil
}

// Price calculates the item price, with discounts
func (c CartItem) GetPrice() float32 {
	return c.Price - c.Discount
}
