package rental_test

import (
	"testing"

	keepertest "rental/testutil/keeper"
	"rental/testutil/nullify"
	rental "rental/x/rental/module"
	"rental/x/rental/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RentalList: []types.Rental{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RentalCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RentalKeeper(t)
	rental.InitGenesis(ctx, k, genesisState)
	got := rental.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RentalList, got.RentalList)
	require.Equal(t, genesisState.RentalCount, got.RentalCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
