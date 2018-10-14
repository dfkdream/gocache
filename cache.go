//Copyright 2018 dfkdream
//License can be found in LICENSE file.

//Package gocache provides LRU based []byte cache.
package gocache

type cacheStorage struct {
	key string
	val []byte
}

//ByteCache is a LRU based []byte cache storage.
type ByteCache struct {
	cacheSize int
	storage   []cacheStorage
}

//NewByteCache returns ByteCache with provided cache size.
func NewByteCache(cacheSize int) *ByteCache {
	return &ByteCache{
		cacheSize: cacheSize,
		storage:   make([]cacheStorage, 0, cacheSize+1),
	}
}

//Set function sets []byte data with provided key string.
//Page replacement will be performed when Set function called.
func (bc *ByteCache) Set(key string, val []byte) {
	for idx, cc := range bc.storage {
		if cc.key == key {
			bc.storage[idx] = cacheStorage{}
			bc.storage = append(bc.storage[:idx], bc.storage[idx+1:]...)
		}
	}
	bc.storage = append([]cacheStorage{cacheStorage{key: key, val: val}}, bc.storage...)
	if len(bc.storage) > bc.cacheSize {
		bc.storage = bc.storage[:bc.cacheSize]
	}
}

//Get function returns cached []byte data with provided key string.
//If there is no byte data with provided key, nil will be returned.
func (bc *ByteCache) Get(key string) []byte {
	for idx, cc := range bc.storage {
		if cc.key == key {
			bc.storage[idx] = cacheStorage{}
			bc.storage = append(bc.storage[:idx], bc.storage[idx+1:]...)
			bc.storage = append([]cacheStorage{cc}, bc.storage...)
			return cc.val
		}
	}
	return nil
}
