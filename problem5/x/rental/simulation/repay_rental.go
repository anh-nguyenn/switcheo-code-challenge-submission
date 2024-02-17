package simulation

import (
	"math/rand"

	"rental/x/rental/keeper"
	"rental/x/rental/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRepayRental(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRepayRental{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RepayRental simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "RepayRental simulation not implemented"), nil, nil
	}
}
