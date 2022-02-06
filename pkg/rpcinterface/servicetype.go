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

	// ServiceTimelord is the timelord service
	ServiceTimelord ServiceType = 5

	// ServicePeer full node service, but for communicating with full nodes using the public protocol
	ServicePeer ServiceType = 6
)

// ServiceType is a type that refers to a particular service
type ServiceType uint8
