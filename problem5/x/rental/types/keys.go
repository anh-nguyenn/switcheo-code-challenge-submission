package types

const (
	// ModuleName defines the module name
	ModuleName = "rental"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rental"
)

var (
	ParamsKey = []byte("p_rental")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	RentalKey      = "Rental/value/"
	RentalCountKey = "Rental/count/"
)
