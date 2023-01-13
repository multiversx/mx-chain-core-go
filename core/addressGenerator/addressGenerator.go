package addressGenerator

import (
	"encoding/binary"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/hashing"
	"github.com/ElrondNetwork/elrond-go-core/hashing/keccak"
)

// addressGenerator is used to generate some addresses based on elrond-go logic
type addressGenerator struct {
	pubkeyConv core.PubkeyConverter
	hasher     hashing.Hasher
}

// NewAddressGenerator will create an address generator instance
func NewAddressGenerator(pubkeyConv core.PubkeyConverter) (*addressGenerator, error) {
	if check.IfNil(pubkeyConv) {
		return nil, core.ErrNilPubkeyConverter
	}

	return &addressGenerator{
		pubkeyConv: pubkeyConv,
		hasher:     keccak.NewKeccak(),
	}, nil
}

// NewAddress is a hook which creates a new smart contract address from the creators address and nonce
// The address is created by applied keccak256 on the appended value off creator address and nonce
// Prefix mask is applied for first 8 bytes 0, and for bytes 9-10 - VM type
// Suffix mask is applied - last 2 bytes are for the shard ID - mask is applied as suffix mask
func (ag *addressGenerator) NewAddress(creatorAddress []byte, creatorNonce uint64, vmType []byte) ([]byte, error) {
	addressLength := ag.pubkeyConv.Len()
	if len(creatorAddress) != addressLength {
		return nil, ErrAddressLengthNotCorrect
	}

	if len(vmType) != core.VMTypeLen {
		return nil, ErrVMTypeLengthIsNotCorrect
	}

	base := hashFromAddressAndNonce(creatorAddress, creatorNonce)
	prefixMask := createPrefixMask(vmType)
	suffixMask := createSuffixMask(creatorAddress)

	copy(base[:core.NumInitCharactersForScAddress], prefixMask)
	copy(base[len(base)-core.ShardIdentiferLen:], suffixMask)

	return base, nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (ag *addressGenerator) IsInterfaceNil() bool {
	return ag == nil
}

func hashFromAddressAndNonce(creatorAddress []byte, creatorNonce uint64) []byte {
	buffNonce := make([]byte, 8)
	binary.LittleEndian.PutUint64(buffNonce, creatorNonce)
	adrAndNonce := append(creatorAddress, buffNonce...)
	scAddress := keccak.NewKeccak().Compute(string(adrAndNonce))

	return scAddress
}

func createPrefixMask(vmType []byte) []byte {
	prefixMask := make([]byte, core.NumInitCharactersForScAddress-core.VMTypeLen)
	prefixMask = append(prefixMask, vmType...)

	return prefixMask
}

func createSuffixMask(creatorAddress []byte) []byte {
	return creatorAddress[len(creatorAddress)-2:]
}
