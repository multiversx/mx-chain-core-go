package slash

import "github.com/ElrondNetwork/elrond-go-core/data"

// SlashingResult contains the slashable data as well as the severity(slashing level)
// for a possible malicious validator
type SlashingResult struct {
	SlashingLevel ThreatLevel
	Headers       []data.HeaderInfoHandler
}
