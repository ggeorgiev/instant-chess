package storage

var cache = initCache()

type Persisted struct {
	size   uint64
	bitset []bool
}
