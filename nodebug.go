// +build !debug,!silent

package ipfscluster

// This is our default logs levels
func init() {
	for _, f := range facilities {
		SetFacilityLogLevel(f, "INFO")
	}

	SetFacilityLogLevel("raft", "ERROR")
	SetFacilityLogLevel("p2p-gorpc", "ERROR")
	//SetFacilityLogLevel("swarm2", l)
	SetFacilityLogLevel("libp2p-raft", "CRITICAL")
}
