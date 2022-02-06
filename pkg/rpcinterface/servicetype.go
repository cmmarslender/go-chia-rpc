package rpcinterface

const (
	// ServiceDaemon the daemon
	ServiceDaemon ServiceType = 0

	// ServiceFullNode the full node service
	ServiceFullNode ServiceType = 1

	// ServiceFarmer the farmer service
	ServiceFarmer ServiceType = 2

	// ServiceHarvester the harvester service
	ServiceHarvester ServiceType = 3

	// ServiceWallet the wallet service
	ServiceWallet ServiceType = 4
)

// ServiceType is a type that refers to a particular service
type ServiceType uint8
