package api

// TokenMetaData is the api metaData struct for tokens
type TokenMetaData struct {
	Nonce      uint64   `json:"nonce"`
	Name       string   `json:"name"`
	Creator    string   `json:"creator"`
	Royalties  uint32   `json:"royalties"`
	Hash       []byte   `json:"hash"`
	URIs       [][]byte `json:"uris"`
	Attributes []byte   `json:"attributes"`
}
