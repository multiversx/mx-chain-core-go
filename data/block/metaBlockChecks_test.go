package block

import (
	"fmt"
	"testing"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var metablockExceptionFields = []string{
	"ShardInfo",
	"PeerInfo",
	"EpochStart",
	"Signature",
	"LeaderSignature",
	"PubKeysBitmap",
	"ReceiptsHash",
	"Reserved",
}

func TestMetaBlockHeader_Checks(t *testing.T) {
	t.Parallel()

	t.Run("nil pointer receiver", func(t *testing.T) {
		t.Parallel()

		var objectToTest *MetaBlock
		err := objectToTest.CheckFieldsForNil()
		require.NotNil(t, err)
		assert.ErrorIs(t, err, data.ErrNilPointerReceiver)
	})
	t.Run("test all fields when set", func(t *testing.T) {
		t.Parallel()

		objectToTest := &MetaBlock{}

		fields := prepareFieldsList(objectToTest, headerV1ExceptionFields...)
		assert.NotEmpty(t, fields)
	})
	t.Run("test all fields when one is unset", func(t *testing.T) {
		t.Parallel()

		objectToTest := &MetaBlock{}

		fields := prepareFieldsList(objectToTest, metablockExceptionFields...)
		assert.NotEmpty(t, fields)

		populateFieldsWithRandomValue(t, objectToTest, fields)
		err := objectToTest.CheckFieldsForNil()
		require.Nil(t, err)

		fmt.Printf("fields tests on %T\n", objectToTest)
		for i := 0; i < len(fields); i++ {
			testField(t, objectToTest, fields, i)
		}
	})
}
