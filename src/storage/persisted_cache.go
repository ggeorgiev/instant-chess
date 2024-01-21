package storage

import lru "github.com/hashicorp/golang-lru/v2"

var persistedCache = initPersistedCache()

func initPersistedCache() *lru.Cache[BoardFigures, ReadMap] {
	lru, err := lru.New[BoardFigures, ReadMap](100000)
	if err != nil {
		panic(err)
	}
	return lru
}

func FetchReadMap(figures BoardFigures) ReadMap {
	if val, ok := persistedCache.Get(figures); ok {
		return val
	}

	return nil
}
