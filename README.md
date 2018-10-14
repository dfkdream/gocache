# gocache
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

import(
    "fmt"
    "strconv"

    "github.com/dfkdream/gocache"
)

func main(){
    cache:=NewByteCache(3)
    for i:=1;i<=5;i++{
        cache.Set(strconv.Itoa(i),[]byte(strconv.Itoa(i*2)))
    }
    for i:=1;i<=5;i++{
        fmt.Println(cache.Get(strconv.Itoa(i)))
    }
}
```