# gocache

[![GoDoc](https://godoc.org/github.com/dfkdream/gocache?status.svg)](https://godoc.org/github.com/dfkdream/gocache)
[![Go Report Card](https://goreportcard.com/badge/github.com/dfkdream/gocache)](https://goreportcard.com/report/github.com/dfkdream/gocache)

gocache is LRU based []byte memory cache library for golang

## Setup

run `go get github.com/dfkdream/gocache`

## Using gocache

* `NewByteCache(cacheSize int)` creates new ByteCache with provided cache size.
Please carefully decide cache size. gocache currently do not calculates total cache size.

* `(bytecache).Set(key string, val []byte)` set []byte data such as image into cache with provided key string.

* `(bytecache).Get(key string)` get []byte data from cache with provided key. If key is not available, `nil` will be returned.

## Example
```go
package main

import (
	"fmt"
	"strconv"

	"github.com/dfkdream/gocache"
)

func main() {
	cache := gocache.NewByteCache(3)
	for i := 1; i <= 5; i++ {
		cache.Set(strconv.Itoa(i), []byte(strconv.Itoa(i*2)))
	}
	for i := 1; i <= 5; i++ {
		if get := cache.Get(strconv.Itoa(i)); get == nil {
			fmt.Println(nil)
		} else {
			fmt.Println(string(get))
		}
	}
}
```
Output
```
<nil>
<nil>
6
8
10
```