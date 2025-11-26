# Toy Sort

## Running

You can build with `make build` and run `./run/toysort`, or just run with `go run cmd/main.go`

## Testing

Unit tests are included, `make test` will run them. `make test-big` will generate a large file and test the sorter on it.

## Benchmark

With a 590mil line file, ~27s
