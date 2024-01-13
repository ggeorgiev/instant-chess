package storage

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	lru "github.com/hashicorp/golang-lru/v2"
)

var cache = initCache()

func initCache() *lru.Cache[[64]peace.Figure, ReadMap] {
	lru, err := lru.New[[64]peace.Figure, ReadMap](100000)
	if err != nil {
		panic(err)
	}
	return lru
}

func Persisted(perm [64]peace.Figure) ReadMap {
	if val, ok := cache.Get(perm); ok {
		return val
	}
	return nil
}
