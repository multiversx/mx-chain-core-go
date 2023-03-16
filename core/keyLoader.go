package core

// keyLoader holds the logic for loading a key from a file and an index
type keyLoader struct {
}

// NewKeyLoader creates a new instance of type key loader
func NewKeyLoader() *keyLoader {
	return &keyLoader{}
}

// LoadKey loads the key with the given index found in the pem file from the given path.
func (kl *keyLoader) LoadKey(path string, skIndex int) ([]byte, string, error) {
	return LoadSkPkFromPemFile(path, skIndex)
}

// LoadAllKeys loads all keys found in the pem file for the given path
func (kl *keyLoader) LoadAllKeys(path string) ([][]byte, []string, error) {
	return LoadAllKeysFromPemFile(path)
}

// IsInterfaceNil returns true if there is no value under the interface
func (kl *keyLoader) IsInterfaceNil() bool {
	return kl == nil
}
