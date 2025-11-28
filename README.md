# Toy Sort

## Running

Requires go 1.23 but would probably work with much older versions.

You can build with `make build` and run `./run/toysort`, or just run with `go run cmd/main.go`

The first line on stdin after running the program should be the path to a file to sort. By defaul `toysort` will print out the top 10 URLs, but you can specify a different count with the `-n` switch.

The format of the input file is expected to be in the form `<url><white_space><long value>`. `toysort` sorts on the 64 bit integer and shows only the URLs in its output.

Example input and output can be found in `testing/small_question.txt` and `testing/small_answer.txt` for `n = 2`

## Testing

Unit tests are included, `make test` will run them. You can generate a large file (default 16mil lines) with `make gen`. And then run `make test-big` to test the sorter on it.

The test works by generating random samples but purposely generating a top ten ahead of time to verify that the program detects it. (An improvement to the test could have these top 10 be randomly distributed throughout the file instead of just at the top)

## Benchmark

With a 500mil line file (~28Gb) on my mac, it took about 360 seconds. 

## Algorithm Assumptions

* Internally uses a heap to store the top N items.
* We assume that the keys are unique (otherwise it will just use the largest example).
* We also ignore lines that don't match the input format.

## Potential enhancements
* The program is built to process the file sequentially (it would work with a stream). Additional speed improvements could be made if we could buffer extremely large files (eg. `mmap`)
* The program is not multi-threaded. You could process several chunks of the file in parallel, and then merge the heaps afterwards.
* Better memory management. Using go `sync.Pool` or something else similar to hold the buffers. You only need to store the top N keys so re-allocating is probably time consuming.
* Should use pproff to find the bottlenecks!

