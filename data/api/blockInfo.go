package api

// BlockInfo is a data transfer object used on the API
type BlockInfo struct {
	Nonce    string `json:"nonce,omitempty"`
	Hash     string `json:"hash,omitempty"`
	RootHash string `json:"RootHash,omitempty"`
}
