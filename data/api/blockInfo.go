package api

// BlockInfo is a data transfer object used on the API
type BlockInfo struct {
	Nonce    uint64 `json:"nonce,omitempty"`
	Hash     string `json:"hash,omitempty"`
	RootHash string `json:"rootHash,omitempty"`
	Epoch    uint32 `json:"epoch,omitempty"`
}
