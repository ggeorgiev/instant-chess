package storage

import lru "github.com/hashicorp/golang-lru/v2"

type FileOffset uint64

var dataBlockCache = initDataBlockCache()

func initDataBlockCache() *lru.Cache[FileOffset, *DataBlock] {
	lru, err := lru.New[FileOffset, *DataBlock](10000)
	if err != nil {
		panic(err)
	}
	return lru
}

func FetchDataBlock(fileOffset FileOffset) *DataBlock {
	if val, ok := dataBlockCache.Get(fileOffset); ok {
		return val
	}

	return nil
}
