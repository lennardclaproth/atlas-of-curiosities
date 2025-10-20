package crypto

import (
	"hash/maphash"
)

func HashBytes(b []byte) uint64 {
	var h maphash.Hash
	h.Write(b)
	return h.Sum64()
}
