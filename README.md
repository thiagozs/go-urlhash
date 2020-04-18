# go-urlhash

Stack of development:

* golang 1.14.1
  
## Build project

Just run `make build` to generate all binaries for distinct OS. Check out on the folder called `out`, inside the folder you find the executable and compressed zip file.

## Running tests

Just run the command `go test ./... -v`

## Running project

To see the commands are suported, just run `go-urlhash -h` for help. 
You see something like that.
```text
➜ go-urlhash.lin -h
Usage of go-urlhash.lin:
  -parallel int
    	Total of number parallel process (default 10)
  -urls value
    	Array of urls for process, you can use with multiple args with -urls x -urls y -urls z OR all urls inside a double quotes, ex: "http://url.com https://niceurl.com ..."
  -version
    	prints current program version
```

## Methods for running

### Default

```text
➜ go-urlhash.lin https://google.com https://twitter.com https://yahoo.com https://adjust.com https://facebook.com
```

### Use the param "parallel"

```text
➜ go-urlhash.lin -parallel 3 https://google.com https://twitter.com https://yahoo.com https://adjust.com https://facebook.com
```

### Use the params "parallel" and "urls"

```text
➜ go-urlhash.lin -parallel 3 -urls https://google.com -urls https://twitter.com -urls https://yahoo.com -urls https://adjust.com -urls https://facebook.com
```

### Use the params "parallel" and "urls" (double quotes)

```text
➜ go-urlhash.lin -parallel 3 -urls "https://google.com https://twitter.com https://yahoo.com https://adjust.com https://facebook.com"
```

## Versioning and license

We use SemVer for versioning. You can see the versions available by checking the tags on this repository.

For more details about our license model, please take a look at the [LICENSE](https://github.com/thiagozs/go-urlhash/blob/master/LICENSE) file.

**2020, thiagozs**