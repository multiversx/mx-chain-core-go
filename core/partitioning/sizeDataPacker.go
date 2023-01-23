package partitioning

import (
	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/batch"
	"github.com/multiversx/mx-chain-core-go/marshal"
)

const minimumMaxPacketSizeInBytes = 1

// SizeDataPacker can split a large slice of byte slices in chunks <= maxPacketSize
// If one element still exceeds maxPacketSize, it will be returned alone
// It does the marshaling of the resulted (smaller) slice of byte slices
type SizeDataPacker struct {
	marshalizer marshal.Marshalizer
}

// NewSizeDataPacker creates a new SizeDataPacker instance
func NewSizeDataPacker(marshalizer marshal.Marshalizer) (*SizeDataPacker, error) {
	if marshalizer == nil || marshalizer.IsInterfaceNil() {
		return nil, core.ErrNilMarshalizer
	}

	return &SizeDataPacker{
		marshalizer: marshalizer,
	}, nil
}

// PackDataInChunks packs the provided data into smaller chunks
// limit is expressed in bytes
func (sdp *SizeDataPacker) PackDataInChunks(data [][]byte, limit int) ([][]byte, error) {
	if limit < minimumMaxPacketSizeInBytes {
		return nil, core.ErrInvalidValue
	}
	if data == nil {
		return nil, core.ErrNilInputData
	}

	returningBuff := make([][]byte, 0)

	elements := make([][]byte, 0)
	lastMarshalized := make([]byte, 0)
	for _, element := range data {
		elements = append(elements, element)
		marshaledElements, err := sdp.marshalizer.Marshal(&batch.Batch{Data: elements})
		if err != nil {
			return nil, err
		}

		isSingleElement := len(elements) == 1
		isMarshaledBuffTooLarge := len(marshaledElements) >= limit

		if isMarshaledBuffTooLarge {
			if isSingleElement {
				returningBuff = append(returningBuff, marshaledElements)
				elements = make([][]byte, 0)
			} else {
				returningBuff = append(returningBuff, lastMarshalized)

				elements = make([][]byte, 0)
				elements = append(elements, element)
				marshaledElements, err = sdp.marshalizer.Marshal(&batch.Batch{Data: elements})
				if err != nil {
					return nil, err
				}

				isMarshaledBuffTooLarge = len(marshaledElements) >= limit
				if isMarshaledBuffTooLarge {
					returningBuff = append(returningBuff, marshaledElements)
					elements = make([][]byte, 0)
				}
			}

			lastMarshalized = make([]byte, 0)
			continue
		}

		lastMarshalized = marshaledElements
	}

	if len(elements) > 0 {
		marshaledElements, err := sdp.marshalizer.Marshal(&batch.Batch{Data: elements})
		if err != nil {
			return nil, err
		}
		returningBuff = append(returningBuff, marshaledElements)
	}

	return returningBuff, nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (sdp *SizeDataPacker) IsInterfaceNil() bool {
	return sdp == nil
}
