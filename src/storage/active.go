package storage

import (
	"encoding/binary"
	"os"
	"path/filepath"
)

const arraySize = 128

type Active struct {
	figures BoardFigures
	size    uint64
	items   uint64
	arrays  map[uint64][arraySize]Data
}

func NewActive(boardFigures BoardFigures, size uint64) *Active {
	return &Active{
		figures: boardFigures,
		size:    size,
		items:   0,
		arrays:  make(map[uint64][arraySize]Data),
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

func (a *Active) Persist(directory string) error {
	filename := a.figures.String()

	fullFilePath := filepath.Join(directory, filename)

	file, err := os.Create(fullFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var sizeBuffer [8]byte // Create an 8-byte buffer
	binary.BigEndian.PutUint64(sizeBuffer[:], a.size)

	_, err = file.Write(sizeBuffer[:])
	if err != nil {
		return err
	}

	var bit []byte

	var bitSizeBuffer [4]byte // Create an 4-byte buffer
	binary.BigEndian.PutUint32(bitSizeBuffer[:], uint32(len(bit)))

	_, err = file.Write(bitSizeBuffer[:])
	if err != nil {
		return err
	}

	return nil
}
