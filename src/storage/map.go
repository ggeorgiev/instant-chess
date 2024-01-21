package storage

import lru "github.com/hashicorp/golang-lru/v2"

type ReadMap interface {
	// Get returns the Data for the given key.
	Get(key BoardStateIndex) *Data
}

type WriteMap interface {
	// Set sets the Data for the given key.
	Set(key BoardStateIndex, data Data)
}

func initCache() *lru.Cache[BoardFigures, ReadMap] {
	lru, err := lru.New[BoardFigures, ReadMap](100000)
	if err != nil {
		panic(err)
	}
	return lru
}

func FetchReadMap(figures BoardFigures) ReadMap {
	if val, ok := cache.Get(figures); ok {
		return val
	}

	return nil
}
