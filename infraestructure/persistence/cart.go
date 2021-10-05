package persistence

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sonalys/ddd/domain/entity"
	"github.com/sonalys/ddd/domain/value"
)

const cartCollection = "cart"

type CartPersistence struct {
	col *mongo.Collection
	ctx context.Context
}

func NewCartPersistence(ctx context.Context, m *Mongo) (*CartPersistence, error) {
	col := m.Collection(cartCollection)
	cart := &CartPersistence{
		col: col,
		ctx: ctx,
	}

	err := cart.createIndexes()
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize cart persistence")
	}

	return cart, nil
}

func (c CartPersistence) createIndexes() error {
	view := c.col.Indexes()

	_, err := view.CreateOne(c.ctx, mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
	})
	if err != nil {
		return errors.Wrap(err, "failed to create indexes")
	}

	return nil
}

func (c CartPersistence) AddItem(ctx context.Context, cartID string, item *entity.CartItem) error {
	cur, err := c.col.UpdateOne(ctx, bson.M{
		"_id": cartID,
	}, bson.M{
		"$push": bson.M{
			"items": item,
		},
	})
	switch err {
	case nil:
	case mongo.ErrNoDocuments:
		return value.ErrNotFound
	default:
		return err
	}

	if cur.MatchedCount != 1 {
		return errors.Errorf("id '%s' affected more %d document", cartID, cur.MatchedCount)
	}
	return err
}

func (c CartPersistence) Create(ctx context.Context) (*entity.Cart, error) {
	cart := &entity.Cart{
		ID: uuid.NewString(),
	}
	_, err := c.col.InsertOne(ctx, cart)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (c CartPersistence) Get(ctx context.Context, cartID string) (*entity.Cart, error) {
	cur := c.col.FindOne(ctx, bson.M{
		"_id": cartID,
	})

	cart := &entity.Cart{}

	err := cur.Decode(cart)
	switch err {
	case nil:
		return cart, nil
	case mongo.ErrNoDocuments:
		return nil, value.ErrNotFound
	default:
		return nil, err
	}
}

func (c CartPersistence) Remove(ctx context.Context, cartID string) error {
	cur, err := c.col.UpdateOne(ctx, bson.M{
		"_id": cartID,
	}, bson.M{
		"updated_at": time.Now(),
	})
	switch err {
	case nil:
	case mongo.ErrNoDocuments:
		return value.ErrNotFound
	default:
		return err
	}

	if cur.MatchedCount != 1 {
		return errors.Errorf("id '%s' affected more %d document", cartID, cur.MatchedCount)
	}
	return err
}

func (c CartPersistence) RemoveItem(ctx context.Context, cartID, itemID string) error {
	cur, err := c.col.UpdateOne(ctx, bson.M{
		"_id": cartID,
	}, bson.M{
		"$pull": bson.M{
			"items.$[].id": itemID,
		},
	})
	switch err {
	case nil:
	case mongo.ErrNoDocuments:
		return value.ErrNotFound
	default:
		return err
	}

	if cur.MatchedCount != 1 {
		return errors.Errorf("id '%s' affected more %d document", cartID, cur.MatchedCount)
	}
	return err
}
