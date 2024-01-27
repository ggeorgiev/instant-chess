package storage

import lru "github.com/hashicorp/golang-lru/v2"

var persistedCache = initPersistedCache()

func initPersistedCache() *lru.Cache[BoardPeaces, ReadMap] {
	lru, err := lru.New[BoardPeaces, ReadMap](100000)
	if err != nil {
		panic(err)
	}
	return lru
}

func FetchReadMap(peaces BoardPeaces) ReadMap {
	if val, ok := persistedCache.Get(peaces); ok {
		return val
	}

	return nil
}
