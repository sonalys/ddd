#!make
include local.env
export

run:
	go run main.go

start_mongo:
	docker run --rm -d -p 27017:27017 --name ddd_mongo mongo:5.0 || true