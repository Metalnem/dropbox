// Package hash implements the Dropbox API content hash algorithm.
package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

// Size is the size of a Dropbox API checksum in bytes.
const Size = 64

// BlockSize is the block size of a Dropbox API checksum in bytes.
const BlockSize = 4 * 1024 * 1024

type digest struct {
	b [][]byte  // hashes of completed blocks
	h hash.Hash // running block hash
	n int       // running block offset
}

// New returns a new hash.Hash computing the Dropbox API checksum.
func New() hash.Hash {
	d := new(digest)
	d.Reset()

	return d
}

func (d *digest) Write(p []byte) (int, error) {
	n := len(p)

	for l := len(p); l > 0; l = len(p) {
		rem := BlockSize - d.n

		if l < rem {
			rem = l
		}

		d.h.Write(p[:rem])
		d.n = (d.n + rem) % BlockSize
		p = p[rem:]

		if d.n == 0 {
			h := d.h.Sum(nil)
			d.b = append(d.b, h)
		}
	}

	return n, nil
}

func (d *digest) Sum(in []byte) []byte {
	h := sha256.New()

	for _, b := range d.b {
		h.Write(hexEncode(b))
	}

	if d.n > 0 {
		h.Write(hexEncode(d.h.Sum(nil)))
	}

	return append(in, hexEncode(h.Sum(nil))...)
}

func (d *digest) Reset() {
	d.b = nil
	d.h = sha256.New()
	d.n = 0
}

func (*digest) Size() int {
	return Size
}

func (*digest) BlockSize() int {
	return BlockSize
}

func hexEncode(b []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)

	return dst
}
