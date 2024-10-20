package types

var (
	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mychain"
)

var (
	ParamsKey = []byte("p_mychain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
