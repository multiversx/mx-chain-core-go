package core_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	A, B int
}

func TestOpenFile_NoExistingFileShouldErr(t *testing.T) {
	t.Parallel()

	file, err := core.OpenFile("testFile1")

	assert.Nil(t, file)
	assert.Error(t, err)
}

func TestOpenFile_NoErrShouldPass(t *testing.T) {
	t.Parallel()

	fileName := "testFile2"
	_, err := os.Create(fileName)
	assert.Nil(t, err)

	file, err := core.OpenFile(fileName)
	if _, errF := os.Stat(fileName); errF == nil {
		_ = os.Remove(fileName)
	}

	assert.NotNil(t, file)
	assert.Nil(t, err)
}

func TestLoadTomlFile_NoExistingFileShouldErr(t *testing.T) {
	t.Parallel()

	cfg := &TestStruct{}

	err := core.LoadTomlFile(cfg, "file")

	assert.Error(t, err)
}

func TestLoadTomlFile_FileExitsShouldPass(t *testing.T) {
	t.Parallel()

	cfg := &TestStruct{}

	fileName := "testFile3"
	_, err := os.Create(fileName)
	assert.Nil(t, err)

	err = core.LoadTomlFile(cfg, fileName)
	if _, errF := os.Stat(fileName); errF == nil {
		_ = os.Remove(fileName)
	}

	assert.Nil(t, err)
}

func TestSaveTomlFile(t *testing.T) {
	t.Parallel()

	t.Run("empty filename, should fail", func(t *testing.T) {
		t.Parallel()

		cfg := &TestStruct{}

		fileName := ""

		err := core.SaveTomlFile(cfg, fileName)
		require.Error(t, err)
	})

	t.Run("invalid path, should fail", func(t *testing.T) {
		t.Parallel()

		cfg := &TestStruct{}

		fileName := "invalid_path" + "/testFile1"

		err := core.SaveTomlFile(cfg, fileName)
		require.Error(t, err)
	})

	t.Run("should work with valid dir", func(t *testing.T) {
		t.Parallel()

		cfg := &TestStruct{}

		dir := t.TempDir()
		fileName := dir + "/testFile1"

		err := core.SaveTomlFile(cfg, fileName)
		require.Nil(t, err)
	})

	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		cfg := &TestStruct{A: 10, B: 20}
		dir := t.TempDir()
		fileName := dir + "/testFile1"

		err := core.SaveTomlFile(cfg, fileName)
		require.Nil(t, err)

		newCfg := &TestStruct{}
		err = core.LoadTomlFile(newCfg, fileName)
		require.Nil(t, err)

		require.Equal(t, cfg, newCfg)
	})
}

func TestLoadJSonFile_NoExistingFileShouldErr(t *testing.T) {
	t.Parallel()

	cfg := &TestStruct{}

	err := core.LoadJsonFile(cfg, "file")

	assert.Error(t, err)
}

func TestLoadJSonFile_FileExitsShouldPass(t *testing.T) {
	t.Parallel()

	cfg := &TestStruct{}

	fileName := "testFile4"
	file, err := os.Create(fileName)
	assert.Nil(t, err)

	data, _ := json.MarshalIndent(TestStruct{A: 0, B: 0}, "", " ")

	_ = ioutil.WriteFile(fileName, data, 0644)

	err = file.Close()
	assert.Nil(t, err)

	err = core.LoadJsonFile(cfg, fileName)
	if _, errF := os.Stat(fileName); errF == nil {
		_ = os.Remove(fileName)
	}

	assert.Nil(t, err)
}

func TestLoadSkPkFromPemFile(t *testing.T) {
	t.Parallel()

	t.Run("invalid index should error", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		dataSk, dataPk, err := core.LoadSkPkFromPemFile(fileName, -1)

		assert.Nil(t, dataSk)
		assert.Empty(t, "", dataPk)
		assert.Equal(t, core.ErrInvalidIndex, err)
	})
	t.Run("missing file should error", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		dataSk, dataPk, err := core.LoadSkPkFromPemFile(fileName, 0)

		assert.Nil(t, dataSk)
		assert.Empty(t, dataPk)
		assert.Error(t, err)
	})
	t.Run("empty file should error", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		_, _ = os.Create(fileName)

		dataSk, dataPk, err := core.LoadSkPkFromPemFile(fileName, 0)

		assert.Nil(t, dataSk)
		assert.Empty(t, dataPk)
		assert.True(t, errors.Is(err, core.ErrEmptyFile))
	})
	t.Run("incorrect header should error", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		file, err := os.Create(fileName)
		assert.Nil(t, err)

		_, _ = file.WriteString("-----BEGIN INCORRECT HEADER ABCD-----\n")
		_, _ = file.WriteString("ChQeKDI8\n")
		_, _ = file.WriteString("-----END INCORRECT HEADER ABCD-----")

		dataSk, dataPk, err := core.LoadSkPkFromPemFile(fileName, 0)

		assert.Nil(t, dataSk)
		assert.Empty(t, dataPk)
		assert.True(t, errors.Is(err, core.ErrPemFileIsInvalid))
	})
	t.Run("invalid pem file should error", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		file, err := os.Create(fileName)
		assert.Nil(t, err)

		_, _ = file.WriteString("data")

		dataSk, dataPk, err := core.LoadSkPkFromPemFile(fileName, 0)

		assert.Nil(t, dataSk)
		assert.Empty(t, dataPk)
		assert.True(t, errors.Is(err, core.ErrPemFileIsInvalid))
	})
	t.Run("invalid index should error", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		file, err := os.Create(fileName)
		assert.Nil(t, err)

		_, _ = file.WriteString("-----BEGIN PRIVATE KEY for data-----\n")
		_, _ = file.WriteString("ChQeKDI8\n")
		_, _ = file.WriteString("-----END PRIVATE KEY for data-----")

		dataSk, dataPk, err := core.LoadSkPkFromPemFile(fileName, 1)

		assert.Nil(t, dataSk)
		assert.Empty(t, dataPk)
		assert.True(t, errors.Is(err, core.ErrInvalidIndex))
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		file, err := os.Create(fileName)
		assert.Nil(t, err)

		skBytes := []byte{10, 20, 30, 40, 50, 60}
		pkString := "ABCD"

		_, _ = file.WriteString("-----BEGIN PRIVATE KEY for " + pkString + "-----\n")
		_, _ = file.WriteString("ChQeKDI8\n")
		_, _ = file.WriteString("-----END PRIVATE KEY for " + pkString + "-----")

		dataSk, dataPk, err := core.LoadSkPkFromPemFile(fileName, 0)

		assert.Equal(t, dataSk, skBytes)
		assert.Equal(t, dataPk, pkString)
		assert.Nil(t, err)
	})
}

func TestLoadAllKeysFromPemFile(t *testing.T) {
	t.Parallel()

	t.Run("missing file should error", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		privateKeys, publicKeys, err := core.LoadAllKeysFromPemFile(fileName)

		assert.Nil(t, privateKeys)
		assert.Nil(t, publicKeys)
		assert.Error(t, err)
	})
	t.Run("empty file should error", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		_, _ = os.Create(fileName)

		privateKeys, publicKeys, err := core.LoadAllKeysFromPemFile(fileName)

		assert.Nil(t, privateKeys)
		assert.Nil(t, publicKeys)
		assert.True(t, errors.Is(err, core.ErrEmptyFile))
	})
	t.Run("incorrect header should error", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		file, err := os.Create(fileName)
		assert.Nil(t, err)

		_, _ = file.WriteString("-----BEGIN INCORRECT HEADER ABCD-----\n")
		_, _ = file.WriteString("ChQeKDI8\n")
		_, _ = file.WriteString("-----END INCORRECT HEADER ABCD-----")

		privateKeys, publicKeys, err := core.LoadAllKeysFromPemFile(fileName)

		assert.Nil(t, privateKeys)
		assert.Empty(t, publicKeys)
		assert.True(t, errors.Is(err, core.ErrPemFileIsInvalid))
	})
	t.Run("invalid pem file should error", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		file, err := os.Create(fileName)
		assert.Nil(t, err)

		_, _ = file.WriteString("data")

		privateKeys, publicKeys, err := core.LoadAllKeysFromPemFile(fileName)

		assert.Nil(t, privateKeys)
		assert.Empty(t, publicKeys)
		assert.True(t, errors.Is(err, core.ErrPemFileIsInvalid))
	})
	t.Run("should work with one key", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		file, err := os.Create(fileName)
		assert.Nil(t, err)

		skBytes := []byte{10, 20, 30, 40, 50, 60}
		pkString := "ABCD"

		_, _ = file.WriteString("-----BEGIN PRIVATE KEY for " + pkString + "-----\n")
		_, _ = file.WriteString("ChQeKDI8\n")
		_, _ = file.WriteString("-----END PRIVATE KEY for " + pkString + "-----")
		_ = file.Close()

		privateKeys, publicKeys, err := core.LoadAllKeysFromPemFile(fileName)

		assert.Equal(t, [][]byte{skBytes}, privateKeys)
		assert.Equal(t, []string{pkString}, publicKeys)
		assert.Nil(t, err)
	})
	t.Run("should work with three keys and extra spaces", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		file, err := os.Create(fileName)
		assert.Nil(t, err)

		skBytes1 := []byte{10, 20, 30, 40, 50, 60}
		pkString1 := "ABCD1"
		skBytes2 := []byte{11, 21, 31, 41, 51, 61}
		pkString2 := "ABCD2"
		skBytes3 := []byte{12, 22, 32, 42, 52, 62}
		pkString3 := "ABCD2"

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

		privateKeys, publicKeys, err := core.LoadAllKeysFromPemFile(fileName)

		assert.Equal(t, [][]byte{skBytes1, skBytes2, skBytes3}, privateKeys)
		assert.Equal(t, []string{pkString1, pkString2, pkString3}, publicKeys)
		assert.Nil(t, err)
	})
}

func TestSaveSkToPemFile(t *testing.T) {
	t.Parallel()

	t.Run("nil file should error", func(t *testing.T) {
		t.Parallel()

		skBytes := make([]byte, 0)
		skBytes = append(skBytes, 10, 20, 30)

		err := core.SaveSkToPemFile(nil, "data", skBytes)

		assert.Equal(t, core.ErrNilFile, err)
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		fileName := filepath.Join(t.TempDir(), "testFile")
		file, err := os.Create(fileName)
		assert.Nil(t, err)

		skBytes := make([]byte, 0)
		skBytes = append(skBytes, 10, 20, 30, 40, 50, 60)

		err = core.SaveSkToPemFile(file, "data", skBytes)
		assert.Nil(t, err)
	})
}

func TestCreateFile(t *testing.T) {
	t.Parallel()

	arg := core.ArgCreateFileArgument{
		Directory:     "subdir",
		Prefix:        "prefix",
		FileExtension: "extension",
	}

	file, err := core.CreateFile(arg)
	assert.Nil(t, err)
	assert.NotNil(t, file)

	assert.True(t, strings.Contains(file.Name(), arg.Prefix))
	assert.True(t, strings.Contains(file.Name(), arg.FileExtension))
	if _, errF := os.Stat(file.Name()); errF == nil {
		_ = os.Remove(file.Name())
		_ = os.Remove(arg.Directory)
	}
}
