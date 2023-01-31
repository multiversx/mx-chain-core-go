package block

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/multiversx/mx-chain-core-go/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var headerV1ExceptionFields = []string{
	"Signature",
	"LeaderSignature",
	"PubKeysBitmap",
	"MetaBlockHashes",
	"EpochStartMetaHash",
	"ReceiptsHash",
	"Reserved",
}

type field struct {
	name          string
	typeValue     string
	objFieldIndex int
}

type fieldsChecker interface {
	CheckFieldsForNil() error
}

func prepareFieldsList(object interface{}, fieldNameExceptions ...string) []field {
	list := make([]field, 0)
	val := reflect.ValueOf(object).Elem()
	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		fieldType := fmt.Sprintf("%v", val.Field(i).Type())
		switch fieldType {
		case "uint64", "uint32", "int", "string":
			continue
		case "block.Type", "[]block.MiniBlockHeader", "[]block.PeerChange":
			continue
		}
		if search(fieldName, fieldNameExceptions...) {
			continue
		}

		list = append(list, field{
			name:          fieldName,
			typeValue:     fieldType,
			objFieldIndex: i,
		})
	}

	return list
}

func search(needle string, haystack ...string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func populateFieldsWithRandomValue(tb testing.TB, object interface{}, fields []field) {
	val := reflect.ValueOf(object)
	for counter, f := range fields {
		fieldValue := val.Elem().FieldByName(f.name)

		switch f.typeValue {
		case "[]uint8":
			fieldValue.SetBytes([]byte(fmt.Sprintf("test field %d", counter)))
		case "[][]uint8":
			fieldValue.Set(reflect.ValueOf([][]byte{
				[]byte(fmt.Sprintf("test field1 %d", counter)),
				[]byte(fmt.Sprintf("test field2 %d", counter)),
			}))
		case "*big.Int":
			fieldValue.Set(reflect.ValueOf(big.NewInt(int64(counter))))
		default:
			assert.Fail(tb, "unimplemented field type "+f.typeValue+" for field "+f.name)
		}
	}
}

func unsetField(tb testing.TB, object interface{}, f field) {
	v := reflect.ValueOf(object)

	fieldValue := v.Elem().FieldByName(f.name)
	switch f.typeValue {
	case "[]uint8", "[][]uint8", "*big.Int":
		fieldValue.Set(reflect.Zero(fieldValue.Type()))
	default:
		assert.Fail(tb, "unimplemented field type "+f.typeValue+" for field "+f.name)
	}
}

func testField(tb testing.TB, object interface{}, fields []field, fieldIndex int) {
	f := fields[fieldIndex]
	fmt.Printf("  testing field %s of type %s\n", f.name, f.typeValue)
	populateFieldsWithRandomValue(tb, object, fields)
	unsetField(tb, object, fields[fieldIndex])

	checker := object.(fieldsChecker)
	err := checker.CheckFieldsForNil()
	require.NotNil(tb, err, "should have return a non nil error for nil field %s", f.name)
	assert.ErrorIs(tb, err, data.ErrNilValue)
}

func TestBlockHeader_Checks(t *testing.T) {
	t.Parallel()

	t.Run("nil pointer receiver", func(t *testing.T) {
		t.Parallel()

		var objectToTest *Header
		err := objectToTest.CheckFieldsForNil()
		require.NotNil(t, err)
		assert.ErrorIs(t, err, data.ErrNilPointerReceiver)
	})
	t.Run("test all fields when set", func(t *testing.T) {
		t.Parallel()

		objectToTest := &Header{}

		fields := prepareFieldsList(objectToTest, headerV1ExceptionFields...)
		assert.NotEmpty(t, fields)
	})
	t.Run("test all fields when one is unset", func(t *testing.T) {
		t.Parallel()

		objectToTest := &Header{}

		fields := prepareFieldsList(objectToTest, headerV1ExceptionFields...)
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
