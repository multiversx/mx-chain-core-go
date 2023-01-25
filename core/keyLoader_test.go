package core

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core/check"
	"github.com/stretchr/testify/assert"
)

func TestNewKeyLoader(t *testing.T) {
	t.Parallel()

	kl := NewKeyLoader()
	assert.False(t, check.IfNil(kl))
}

func TestKeyLoader_LoadKey(t *testing.T) {
	t.Parallel()

	kl := NewKeyLoader()
	file, sourceData := createTestFile(t)
	recoveredData := make(map[string][]byte)
	for i := 0; i < 3; i++ {
		sk, pk, err := kl.LoadKey(file, i)
		assert.Nil(t, err)
		recoveredData[pk] = sk
	}

	assert.Equal(t, sourceData, recoveredData)
}

func TestKeyLoader_LoadAllKeys(t *testing.T) {
	t.Parallel()

	kl := NewKeyLoader()
	file, sourceData := createTestFile(t)

	recoveredData := make(map[string][]byte)
	privateKeys, publicKeys, err := kl.LoadAllKeys(file)
	assert.Nil(t, err)

	for i, pk := range publicKeys {
		recoveredData[pk] = privateKeys[i]
	}

	assert.Equal(t, sourceData, recoveredData)
}

func createTestFile(tb testing.TB) (string, map[string][]byte) {
	fileName := filepath.Join(tb.TempDir(), "testFile")
	file, err := os.Create(fileName)
	assert.Nil(tb, err)

	pkString1 := "ABCD1"
	pkString2 := "ABCD2"
	pkString3 := "ABCD2"

	data := make(map[string][]byte)
	data[pkString1] = []byte{10, 20, 30, 40, 50, 60}
	data[pkString2] = []byte{11, 21, 31, 41, 51, 61}
	data[pkString3] = []byte{12, 22, 32, 42, 52, 62}

	_, _ = file.WriteString("-----BEGIN PRIVATE KEY for " + pkString1 + "-----\n")
	_, _ = file.WriteString("ChQeKDI8\n")
	_, _ = file.WriteString("-----END PRIVATE KEY for " + pkString1 + "-----\n")
	_, _ = file.WriteString("-----BEGIN PRIVATE KEY for " + pkString2 + "-----\n\n\n")
	_, _ = file.WriteString("CxUfKTM9\n")
	_, _ = file.WriteString("-----END PRIVATE KEY for " + pkString2 + "-----\n")
	_, _ = file.WriteString("-----BEGIN PRIVATE KEY for " + pkString3 + "-----\n")
	_, _ = file.WriteString("DBYgKjQ+\n")
	_, _ = file.WriteString("-----END PRIVATE KEY for " + pkString3 + "-----\n\n\n\n")
	_ = file.Close()

	return fileName, data
}
