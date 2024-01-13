package storage

type ReadMap interface {
	// Get returns the Data for the given key.
	Get(key uint64) Data
}

type WriteMap interface {
	// Set sets the Data for the given key.
	Set(key uint64, data Data)
}
