package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RentalList: []Rental{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in rental
	rentalIdMap := make(map[uint64]bool)
	rentalCount := gs.GetRentalCount()
	for _, elem := range gs.RentalList {
		if _, ok := rentalIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for rental")
		}
		if elem.Id >= rentalCount {
			return fmt.Errorf("rental id should be lower or equal than the last id")
		}
		rentalIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
