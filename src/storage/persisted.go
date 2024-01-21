package storage

const DataBlockSize = 4096

type DataBlock [DataBlockSize]Data

type Persisted struct {
	size   uint64
	bitset []bool
}
