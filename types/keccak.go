package types

import "hash"

type KeccakState interface {
	hash.Hash
	Read([]byte) (int, error)
}

