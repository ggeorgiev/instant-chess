package storage

const arraySize = 128

type Active struct {
	items  uint64
	arrays map[uint64][arraySize]Data
}

func NewActive() *Active {
	return &Active{
		items:  0,
		arrays: make(map[uint64][arraySize]Data),
	}
}

func (a *Active) Get(key uint64) Data {
	return a.arrays[key/arraySize][a.items%arraySize]
}

func (a *Active) Set(key uint64, data Data) {
	blockKey := key / arraySize
	if array, ok := a.arrays[blockKey]; ok {
		array[a.items%arraySize] = data
	} else {
		var array [arraySize]Data
		array[a.items%arraySize] = data
		a.arrays[blockKey] = array
	}
	a.items++
}

func (a *Active) Ratio() float64 {
	return float64(a.items) / float64(len(a.arrays)*arraySize)
}
