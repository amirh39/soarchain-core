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
	EpochDataKey = "EpochData-value-"
)

const (
	MasterKeyKey = "MasterKey-value-"
)

const (
	FactoryKeysKey      = "FactoryKeys-value-"
	FactoryKeysCountKey = "FactoryKeys-count-"
)

var (
	ProposalsKeyPrefix            = []byte{0x00}
	ActiveProposalQueuePrefix     = []byte{0x01}
	InactiveProposalQueuePrefix   = []byte{0x02}
	ProposalIDKey                 = []byte{0x03}
	VotingPeriodProposalKeyPrefix = []byte{0x04}

	BalancesPrefix = []byte{0x02}
	FeePoolKey     = []byte{0x00}

	DepositsKeyPrefix = []byte{0x10}

	VotesKeyPrefix = []byte{0x20}

	// ParamsKey is the key to query all gov params
	ParamsKey = []byte{0x30}
)
