package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/sonalys/ddd/application"
	"github.com/sonalys/ddd/infraestructure/persistence"
	"github.com/sonalys/ddd/interfaces"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// INFO: stop channel is needed for graceful shutdown implementation.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	mongo, err := persistence.NewMongoClient(ctx, persistence.MongoConfig{
		Hosts:    []string{":27500"},
		Username: "123",
		Password: "456",
	})
	if err != nil {
		panic(err)
	}

	cart, err := application.NewCartApp(ctx, mongo)
	if err != nil {
		panic(err)
	}

	handler := interfaces.NewRouter(interfaces.Dependency{
		Cart: cart,
	})

	go func() {
		http.ListenAndServe(":8080", handler)
	}()
	<-stop
	cancel()
}
