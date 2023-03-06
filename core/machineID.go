package core

import "github.com/denisbrodbeck/machineid"

// GetAnonymizedMachineID returns the machine ID anonymized with the provided app ID string
func GetAnonymizedMachineID(appID string) string {
	machineID, err := machineid.ProtectedID(appID)
	if err != nil {
		machineID = "unknown machine ID"
	}
	if len(machineID) > MaxMachineIDLen {
		machineID = machineID[:MaxMachineIDLen]
	}

	return machineID
}
