package keyValStorage

import "github.com/multiversx/mx-chain-core-go/core"

// KeyValStorage holds a key and an associated value
type keyValStorage struct {
	key     []byte
	value   []byte
	version core.TrieNodeVersion
}

// NewKeyValStorage creates a new key-value storage
func NewKeyValStorage(key []byte, val []byte, version core.TrieNodeVersion) *keyValStorage {
	return &keyValStorage{
		key:     key,
		value:   val,
		version: version,
	}
}

// Key returns the key in the key-value storage
func (k *keyValStorage) Key() []byte {
	return k.key
}

// Value returns the value in the key-value storage
func (k *keyValStorage) Value() []byte {
	return k.value
}

// Version returns the version in the key-value storage
func (k *keyValStorage) Version() core.TrieNodeVersion {
	return k.version
}
