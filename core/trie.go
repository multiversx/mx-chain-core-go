package core

import (
	"strconv"

	"github.com/multiversx/mx-chain-core-go/core/check"
)

// TrieNodeVersion defines the version of the trie node
type TrieNodeVersion uint8

const (
	// NotSpecified means that the value is not populated or is not important
	NotSpecified TrieNodeVersion = iota

	// AutoBalanceEnabled is used for data tries, and only after the activation of AutoBalanceDataTriesEnableEpoch flag
	AutoBalanceEnabled
)

const (
	// NotSpecifiedString is the string representation of NotSpecified trie node version
	NotSpecifiedString = "not specified"

	// AutoBalanceEnabledString is the string representation of AutoBalanceEnabled trie node version
	AutoBalanceEnabledString = "auto balanced"
)

func (version TrieNodeVersion) String() string {
	switch version {
	case NotSpecified:
		return NotSpecifiedString
	case AutoBalanceEnabled:
		return AutoBalanceEnabledString
	default:
		return "unknown: " + strconv.Itoa(int(version))
	}
}

type trieNodeVersionVerifier struct {
	enableEpochsHandler EnableEpochsHandler
}

func NewTrieNodeVersionVerifier(enableEpochsHandler EnableEpochsHandler) (*trieNodeVersionVerifier, error) {
	if check.IfNil(enableEpochsHandler) {
		return nil, ErrNilEnableEpochsHandler
	}

	return &trieNodeVersionVerifier{
		enableEpochsHandler: enableEpochsHandler,
	}, nil
}

// IsValidVersion returns true if the given trie node version is valid
func (vv *trieNodeVersionVerifier) IsValidVersion(version TrieNodeVersion) bool {
	if vv.enableEpochsHandler.IsAutoBalanceDataTriesEnabled() {
		return version <= AutoBalanceEnabled
	}

	return version == NotSpecified
}

// IsInterfaceNil returns true if there is no value under the interface
func (vv *trieNodeVersionVerifier) IsInterfaceNil() bool {
	return vv == nil
}

// GetVersionForNewData returns the trie node version that should be used for new data
func GetVersionForNewData(handler EnableEpochsHandler) TrieNodeVersion {
	if handler.IsAutoBalanceDataTriesEnabled() {
		return AutoBalanceEnabled
	}

	return NotSpecified
}

// TrieData holds the data that will be inserted into the trie
type TrieData struct {
	Key     []byte
	Value   []byte
	Version TrieNodeVersion
}
