package core

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadTomlFileShouldFailBecauseOfMissingField(t *testing.T) {
	var epochsConfig testConfig

	err := LoadTomlFileWithDefaultChecks(&epochsConfig, "./testdata/incompleteTomlFile.toml")
	require.Error(t, err)
	require.Contains(t, err.Error(), "config value for field Efg not set")
}

func TestLoadTomlFileWithDefaultChecks(t *testing.T) {
	var epochsConfig testConfig

	err := LoadTomlFileWithDefaultChecks(&epochsConfig, "./testdata/okTomlFile.toml")
	require.NoError(t, err)
}

type testConfig struct {
	FirstConfigSection  firstConfigSection
	SecondConfigSection secondConfigSection
}

type firstConfigSection struct {
	Abc         uint32
	Bcd         uint32
	InnerConfig testInnerConfigSection
	Cde         string
	Ghi         bool
}

type testInnerConfigSection struct {
	Def string
	Efg uint32
}

type secondConfigSection struct {
	Fgh uint32
}
