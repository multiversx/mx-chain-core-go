package api

// Guardian holds the relevant information for an account guardian
type Guardian struct {
	Address         string `json:"address"`
	ActivationEpoch uint32 `json:"activationEpoch"`
	ServiceUID      string `json:"serviceUID"`
}

// GuardianData holds data relating to the configured guardian(s) and guarded state of an account
type GuardianData struct {
	ActiveGuardian  *Guardian `json:"activeGuardian,omitempty"`
	PendingGuardian *Guardian `json:"pendingGuardian,omitempty"`
	Guarded         bool      `json:"guarded"`
}
