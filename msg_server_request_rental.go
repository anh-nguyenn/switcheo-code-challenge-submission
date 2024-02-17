package keeper

import (
	"context"

	"rental/x/rental/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Define the minimum amount threshold as a constant
const MinimumAmountThreshold = 1000000

func (k msgServer) RequestRental(goCtx context.Context, msg *types.MsgRequestRental) (*types.MsgRequestRentalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid amount")
	}

	// Check if the requested amount meets the minimum threshold
	if amount.IsAnyLT(sdk.NewCoins(sdk.NewInt64Coin(amount.GetDenomByIndex(0), MinimumAmountThreshold))) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "the requested amount is less than the minimum threshold of %d", MinimumAmountThreshold)
	}
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
