package types

const (
	// ModuleName defines the module name
	ModuleName = "poa"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_poa"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	TotalClientsKey = "TotalClients-value-"
)

const (
	TotalChallengersKey = "TotalChallengers-value-"
)

const (
	TotalRunnersKey = "TotalRunners-value-"
)
