package storage

type ReadMap interface {
	// Get returns the Data for the given key.
	Get(key BoardStateIndex) *Data
}

type WriteMap interface {
	// Set sets the Data for the given key.
	Set(key BoardStateIndex, data Data)
}
