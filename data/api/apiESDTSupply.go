package api

// ESDTSupply represents the structure for esdt supply that is returned by api routes
type ESDTSupply struct {
	InitialMinted    string `json:"initialMinted"`
	Supply           string `json:"supply"`
	Burned           string `json:"burned"`
	Minted           string `json:"minted"`
	RecomputedSupply bool   `json:"recomputedSupply"`
}
