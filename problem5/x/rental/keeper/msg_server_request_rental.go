package keeper

import (
	"context"

	"rental/x/rental/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RequestRental(goCtx context.Context, msg *types.MsgRequestRental) (*types.MsgRequestRentalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	var rental = types.Rental{
		Amount:  msg.Amount,
		Fee:     msg.Fee,
		Asset:   msg.Asset,
		DueDate: msg.DueDate,
		State:   "requested",
		Renter:  msg.Creator,
	}
	renter, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	asset, err := sdk.ParseCoinsNormalized(rental.Asset)
	if err != nil {
		panic(err)
	}
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, renter, types.ModuleName, asset)
	if sdkError != nil {
		return nil, sdkError
	}
	k.AppendRental(ctx, rental)

	return &types.MsgRequestRentalResponse{}, nil
}
