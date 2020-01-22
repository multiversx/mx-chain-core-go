package hashing_test

import (
	"testing"

	"github.com/ElrondNetwork/elrond-go/hashing"
	"github.com/ElrondNetwork/elrond-go/hashing/blake2b"
	"github.com/ElrondNetwork/elrond-go/hashing/fnv"
	"github.com/ElrondNetwork/elrond-go/hashing/keccak"
	"github.com/ElrondNetwork/elrond-go/hashing/sha256"
	"github.com/stretchr/testify/assert"
)

func TestSha256(t *testing.T) {
	Suite(t, sha256.Sha256{})
}

func TestBlake2b(t *testing.T) {
	Suite(t, blake2b.Blake2b{})
}

func TestKeccak(t *testing.T) {
	Suite(t, keccak.Keccak{})
}

func TestFnv(t *testing.T) {
	Suite(t, fnv.Fnv{})
}

func Suite(t *testing.T, h hashing.Hasher) {
	TestingNilInterface(t, h)
	TestingSize(t, h)
	TestingCalculateHash(t, h)
	TestingCalculateEmptyHash(t, h)
	TestingNilReturn(t, h)

}

func TestingNilInterface(t *testing.T, h hashing.Hasher) {

	res := h.IsInterfaceNil()

	assert.False(t, res)

}

func TestingSize(t *testing.T, h hashing.Hasher) {

	res := h.Size()

	assert.Equal(t, 0, res%2)

}

func TestingCalculateHash(t *testing.T, h hashing.Hasher) {

	h1 := h.Compute("a")
	h2 := h.Compute("b")

	assert.NotEqual(t, h1, h2)

}

func TestingCalculateEmptyHash(t *testing.T, h hashing.Hasher) {
	h1 := h.Compute("")
	h2 := h.EmptyHash()

	assert.Equal(t, h1, h2)

}

func TestingNilReturn(t *testing.T, h hashing.Hasher) {
	h1 := h.Compute("a")
	assert.NotNil(t, h1)
}
