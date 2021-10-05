package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Cart represents a cart, which can hold multiple items.
type Cart struct {
	ID                string      `json:"id" bson:"_id"`
	Items             []*CartItem `json:"items" bson:"items"`
	TransportationFee float32     `json:"transportation_fee" bson:"transportationFee"`

	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	DeletedAt time.Time `bson:"deletedAt"`
}

var (
	// ErrInvalidCart indicates a cart contains errors.
	ErrInvalidCart error = errors.New("cart is invalid")
)

// Valid implements Validatable interface.
func (c Cart) Valid() error {
	cartErrors := []string{}

	if c.TransportationFee < 0 {
		cartErrors = append(cartErrors, "transportation fee must be greater than 0")
	}

	for i, cartItem := range c.Items {
		if err := cartItem.Valid(); err != nil {
			msg := fmt.Sprintf("item %d: %s", i, err)
			cartErrors = append(cartErrors, msg)
		}
	}

	if len(cartErrors) > 0 {
		msg := strings.Join(cartErrors, ";")
		return errors.Wrap(ErrInvalidCart, msg)
	}
	return nil
}

// GetValues returns the sum of all prices and discounts in this cart.
func (c Cart) GetValues() (priceSum, discountSum float32) {
	for _, item := range c.Items {
		priceSum += item.Price
		discountSum += item.Discount
	}
	return
}
