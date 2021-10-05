package interfaces

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/sonalys/ddd/application"
	"github.com/sonalys/ddd/domain/entity"
	"github.com/sonalys/ddd/domain/value"
)

// Decoder is a cache for schema decoding.
var decoder = schema.NewDecoder()

type cartHandler struct {
	application.Cart
}

func newCartHandler(cart application.Cart, r *mux.Router) cartHandler {
	h := cartHandler{
		Cart: cart,
	}

	r.Path("/").
		Methods(http.MethodPost).
		HandlerFunc(h.create)

	r.Path("/").
		Methods(http.MethodDelete).
		HandlerFunc(h.remove)

	r.Path("/").
		Methods(http.MethodGet).
		HandlerFunc(h.get)

	r.Path("/items").
		Methods(http.MethodPost).
		HandlerFunc(h.addItem)

	r.Path("/items").
		Methods(http.MethodDelete).
		HandlerFunc(h.remove)

	return h
}

func (c *cartHandler) create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	cart, err := c.Cart.Create(ctx)
	if err != nil {
		httpErr := newProcessError(err)
		writeJSON(w, httpErr, http.StatusInternalServerError)
		return
	}

	err = writeJSON(w, cart, http.StatusCreated)
	if err != nil {
		httpErr := newEncodeError(err)
		writeJSON(w, httpErr, http.StatusInternalServerError)
	}
}

// CartIDRequest is a contract for requests with cart_id in the query.
type CartIDRequest struct {
	CartID string `schema:"cart_id"`
}

// Decode implements Decodable interface.
func (b *CartIDRequest) Decode(r *http.Request) error {
	values := r.URL.Query()

	err := decoder.Decode(b, values)
	if err != nil {
		return err
	}

	return nil
}

func (b CartIDRequest) Valid() error {
	if b.CartID == "" {
		return errors.New("cart_id cannot be empty")
	}

	return nil
}

func (c *cartHandler) get(w http.ResponseWriter, r *http.Request) {
	var req CartIDRequest
	ctx := r.Context()

	if err := parseRequest(r, &req); err != nil {
		httpErr := newRequestError(err)
		writeJSON(w, httpErr, http.StatusBadRequest)
		return
	}

	cart, err := c.Cart.Get(ctx, req.CartID)
	switch err {
	case nil:
		writeJSON(w, cart, http.StatusOK)
	case value.ErrNotFound:
		httpErr := newNotFoundError(err)
		writeJSON(w, httpErr, http.StatusNotFound)
	default:
		httpErr := newProcessError(err)
		writeJSON(w, httpErr, http.StatusInternalServerError)
	}
}

func (c *cartHandler) remove(w http.ResponseWriter, r *http.Request) {
	var req CartIDRequest
	ctx := r.Context()

	if err := parseRequest(r, &req); err != nil {
		httpErr := newRequestError(err)
		writeJSON(w, httpErr, http.StatusBadRequest)
		return
	}

	err := c.Cart.Remove(ctx, req.CartID)
	switch err {
	case nil:
		w.WriteHeader(http.StatusNoContent)
	case value.ErrNotFound:
		httpErr := newNotFoundError(err)
		writeJSON(w, httpErr, http.StatusNotFound)
	default:
		httpErr := newProcessError(err)
		writeJSON(w, httpErr, http.StatusInternalServerError)
	}
}

// AddCartItemRequest is a contract for adding cart items.
type AddCartItemRequest struct {
	CartID   string `schema:"cart_id"`
	CartItem entity.CartItem
}

// Decode implements Decodable interface.
func (b *AddCartItemRequest) Decode(r *http.Request) error {
	values := r.URL.Query()

	err := decoder.Decode(b, values)
	if err != nil {
		return err
	}

	json.NewDecoder(r.Body).Decode(&b.CartItem)
	if err != nil {
		return err
	}

	return nil
}

func (b AddCartItemRequest) Valid() error {
	if b.CartID == "" {
		return errors.New("cart_id cannot be empty")
	}

	if err := b.CartItem.Valid(); err != nil {
		return err
	}

	return nil
}

func (c *cartHandler) addItem(w http.ResponseWriter, r *http.Request) {
	var req AddCartItemRequest
	ctx := r.Context()

	if err := parseRequest(r, &req); err != nil {
		httpErr := newRequestError(err)
		writeJSON(w, httpErr, http.StatusBadRequest)
		return
	}

	err := c.Cart.AddItem(ctx, req.CartID, &req.CartItem)
	switch err {
	case nil:
		w.WriteHeader(http.StatusNoContent)
	case value.ErrNotFound:
		httpErr := newNotFoundError(err)
		writeJSON(w, httpErr, http.StatusNotFound)
	default:
		httpErr := newProcessError(err)
		writeJSON(w, httpErr, http.StatusInternalServerError)
	}
}

// RemoveCartItemRequest is a contract for adding cart items.
type RemoveCartItemRequest struct {
	CartID     string `schema:"cart_id"`
	CartItemID string `schema:"cart_item_id"`
}

// Decode implements Decodable interface.
func (b *RemoveCartItemRequest) Decode(r *http.Request) error {
	values := r.URL.Query()

	err := decoder.Decode(b, values)
	if err != nil {
		return err
	}

	return nil
}

func (b RemoveCartItemRequest) Valid() error {
	if b.CartID == "" {
		return errors.New("cart_id cannot be empty")
	}

	if b.CartItemID == "" {
		return errors.New("cart_id cannot be empty")
	}
	return nil
}

func (c *cartHandler) removeItem(w http.ResponseWriter, r *http.Request) {
	var req RemoveCartItemRequest
	ctx := r.Context()

	if err := parseRequest(r, &req); err != nil {
		httpErr := newRequestError(err)
		writeJSON(w, httpErr, http.StatusBadRequest)
		return
	}

	err := c.Cart.RemoveItem(ctx, req.CartID, req.CartItemID)
	switch err {
	case nil:
		w.WriteHeader(http.StatusNoContent)
	case value.ErrNotFound:
		httpErr := newNotFoundError(err)
		writeJSON(w, httpErr, http.StatusNotFound)
	default:
		httpErr := newProcessError(err)
		writeJSON(w, httpErr, http.StatusInternalServerError)
	}
}
