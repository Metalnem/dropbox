// Package hash implements the Dropbox API content hash algorithm.
package hash

import "hash"

// Size is the size of a Dropbox API checksum in bytes.
const Size = 64

// BlockSize is the block size of a Dropbox API checksum in bytes.
const BlockSize = 4 * 1024 * 1024

type digest struct {
}

// New returns a new hash.Hash computing the Dropbox API checksum.
func New() hash.Hash {
	return new(digest)
}

func (*digest) Write(p []byte) (int, error) {
	return 0, nil
}

func (*digest) Sum(b []byte) []byte {
	return nil
}

func (*digest) Reset() {
}

func (*digest) Size() int {
	return Size
}

func (*digest) BlockSize() int {
	return BlockSize
}
