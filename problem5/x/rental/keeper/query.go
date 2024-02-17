package keeper

import (
	"rental/x/rental/types"
)

var _ types.QueryServer = Keeper{}
