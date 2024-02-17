package keeper_test

import (
	"context"
	"testing"

	keepertest "rental/testutil/keeper"
	"rental/testutil/nullify"
	"rental/x/rental/keeper"
	"rental/x/rental/types"

	"github.com/stretchr/testify/require"
)

func createNRental(keeper keeper.Keeper, ctx context.Context, n int) []types.Rental {
	items := make([]types.Rental, n)
	for i := range items {
		items[i].Id = keeper.AppendRental(ctx, items[i])
	}
	return items
}

func TestRentalGet(t *testing.T) {
	keeper, ctx := keepertest.RentalKeeper(t)
	items := createNRental(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRental(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestRentalRemove(t *testing.T) {
	keeper, ctx := keepertest.RentalKeeper(t)
	items := createNRental(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRental(ctx, item.Id)
		_, found := keeper.GetRental(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRentalGetAll(t *testing.T) {
	keeper, ctx := keepertest.RentalKeeper(t)
	items := createNRental(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRental(ctx)),
	)
}

func TestRentalCount(t *testing.T) {
	keeper, ctx := keepertest.RentalKeeper(t)
	items := createNRental(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRentalCount(ctx))
}
