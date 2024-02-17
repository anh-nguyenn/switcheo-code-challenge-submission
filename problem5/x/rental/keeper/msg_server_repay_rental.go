package keeper

import (
	"context"

	"rental/x/rental/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RepayRental(goCtx context.Context, msg *types.MsgRepayRental) (*types.MsgRepayRentalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	rental, found := k.GetRental(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if rental.State != "approved" {
		return nil, errorsmod.Wrapf(types.ErrWrongRentalState, "%v", rental.State)
	}
	renter, _ := sdk.AccAddressFromBech32(rental.Renter)
	owner, _ := sdk.AccAddressFromBech32(rental.Owner)
	if msg.Creator != rental.Renter {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot repay: not the renter")
	}
	amount, _ := sdk.ParseCoinsNormalized(rental.Amount)
	fee, _ := sdk.ParseCoinsNormalized(rental.Fee)
	asset, _ := sdk.ParseCoinsNormalized(rental.Asset)
	err := k.bankKeeper.SendCoins(ctx, renter, owner, amount)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoins(ctx, renter, owner, fee)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, renter, asset)
	if err != nil {
		return nil, err
	}
	rental.State = "repayed"
	k.SetRental(ctx, rental)
	return &types.MsgRepayRentalResponse{}, nil
}
