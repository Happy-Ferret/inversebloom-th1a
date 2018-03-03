/*
	Inverse Bloom InverseBloomFilter (aka Opposite Of A Bloom InverseBloomFilter)
	- https://github.com/jmhodges/opposite_of_a_bloom_filter

	modified to use th1a hash algorithm
*/
package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"sync/atomic"
	"unsafe"
)

var (
	maxInverseBloomFilterSize = 1 << 30
)

// InverseBloomFilter is an inverse bloom filter.
type InverseBloomFilter struct {
	array    []*[]byte
	sizeMask uint32
	*TH1A
}

// NewInverseBloomFilter returns a new Inverse Bloom filter of provided size.
func NewInverseBloomFilter(size int) (*InverseBloomFilter, error) {
	if size > maxInverseBloomFilterSize {
		return nil, errors.New("size too large to round to a power of 2")
	}
	if size <= 0 {
		return nil, errors.New("filter cannot have zero or negative size")
	}
	// round to the next largest power of two
	size = int(math.Pow(2, math.Ceil(math.Log2(float64(size)))))
	slice := make([]*[]byte, size)
	sizeMask := uint32(size - 1)
	tt := &TH1A{0x0102030405060788}
	return &InverseBloomFilter{slice, sizeMask, tt}, nil
}

// Contains returns true of the Inverse Bloom filter contains the given id.
func (f *InverseBloomFilter) Contains(id []byte) bool {
	s64 := f.Sum64(id)
	s64b := make([]byte, 8)
	binary.LittleEndian.PutUint64(s64b, s64)

	s32 := uint32(s64b[0])
	for _, val := range s64b[1:3] {
		s32 = s32 << 3
		s32 += uint32(val)
	}

	uindex := s32 & f.sizeMask
	index := int32(uindex)
	oldId := getAndSet(f.array, index, id)
	return bytes.Equal(oldId, id)
}

// Size returns the size of the inverse bloom filter.
func (f *InverseBloomFilter) Size() int {
	return len(f.array)
}

// Returns the id that was in the slice at the given index after putting the
// new id in the slice at that index, atomically.
func getAndSet(arr []*[]byte, index int32, id []byte) []byte {
	indexPtr := (*unsafe.Pointer)(unsafe.Pointer(&arr[index]))
	idUnsafe := unsafe.Pointer(&id)
	var oldId []byte
	for {
		oldIdUnsafe := atomic.LoadPointer(indexPtr)
		if atomic.CompareAndSwapPointer(indexPtr, oldIdUnsafe, idUnsafe) {
			oldIdPtr := (*[]byte)(oldIdUnsafe)
			if oldIdPtr != nil {
				oldId = *oldIdPtr
			}
			break
		}
	}
	return oldId
}
