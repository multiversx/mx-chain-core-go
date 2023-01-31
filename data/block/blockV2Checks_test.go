package block

import (
	"fmt"
	"testing"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var headerV2ExceptionFields = []string{
	"Header",
	"ScheduledRootHash",
}

func TestBlockHeaderV2_Checks(t *testing.T) {
	t.Parallel()

	t.Run("nil pointer receiver", func(t *testing.T) {
		t.Parallel()

		var objectToTest *HeaderV2
		err := objectToTest.CheckFieldsForNil()
		require.NotNil(t, err)
		assert.ErrorIs(t, err, data.ErrNilPointerReceiver)
	})
	t.Run("inner header is a nil pointer receiver", func(t *testing.T) {
		t.Parallel()

		objectToTest := &HeaderV2{}
		err := objectToTest.CheckFieldsForNil()
		require.NotNil(t, err)
		assert.ErrorIs(t, err, data.ErrNilPointerReceiver)
	})
	t.Run("test all fields when set", func(t *testing.T) {
		t.Parallel()

		objectToTest := &HeaderV2{}

		fields := prepareFieldsList(objectToTest, headerV1ExceptionFields...)
		assert.NotEmpty(t, fields)
	})
	t.Run("test all fields when one is unset on inner Header", func(t *testing.T) {
		t.Parallel()

		objectToTest := &HeaderV2{
			Header: &Header{},
		}

		fieldsForHeaderV2 := prepareFieldsList(objectToTest, headerV2ExceptionFields...)
		assert.NotEmpty(t, fieldsForHeaderV2)
		populateFieldsWithRandomValue(t, objectToTest, fieldsForHeaderV2)

		fieldsForHeaderV1 := prepareFieldsList(objectToTest.Header, headerV1ExceptionFields...)
		assert.NotEmpty(t, fieldsForHeaderV1)

		populateFieldsWithRandomValue(t, objectToTest.Header, fieldsForHeaderV1)
		err := objectToTest.CheckFieldsForNil()
		require.Nil(t, err)

		fmt.Printf("fields tests on %T\n", objectToTest.Header)
		for i := 0; i < len(fieldsForHeaderV1); i++ {
			testField(t, objectToTest.Header, fieldsForHeaderV1, i)
		}
	})
	t.Run("test all fields when one is unset on HeaderV2", func(t *testing.T) {
		t.Parallel()

		objectToTest := &HeaderV2{
			Header: &Header{},
		}

		fieldsForHeaderV1 := prepareFieldsList(objectToTest.Header, headerV1ExceptionFields...)
		assert.NotEmpty(t, fieldsForHeaderV1)
		populateFieldsWithRandomValue(t, objectToTest.Header, fieldsForHeaderV1)

		fields := prepareFieldsList(objectToTest, headerV2ExceptionFields...)
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
