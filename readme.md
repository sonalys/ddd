# Domain Driven Design Cart Api

This is a POC concept for testing domain driven design ( DDD )
The main objective of this POC is to learn and validate all the benefits which DDD provides for code architecture.

## Running the project

The project uses `local.env` variables to setup external dependencies, also for configuring the http listener, which defaults to address `:8080`

To start the api, simple run `make start`

## API Specifications

This API follows OPEN API specifications, to visualize all the contracts, run `make serve_swagger`

## Testing

To run the tests, you can either:

- `make test` to run tests without coverage
- `make coverage` to get code coverage
