package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "rental/testutil/keeper"
	"rental/x/rental/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.RentalKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
