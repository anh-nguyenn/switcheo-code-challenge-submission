package keeper

import (
	"context"

	"rental/x/rental/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CancelRental(goCtx context.Context, msg *types.MsgCancelRental) (*types.MsgCancelRentalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	rental, found := k.GetRental(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if rental.Renter != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot cancel: not the renter")
	}
	if rental.State != "requested" {
		return nil, errorsmod.Wrapf(types.ErrWrongRentalState, "%v", rental.State)
	}
	renter, _ := sdk.AccAddressFromBech32(rental.Renter)
	asset, _ := sdk.ParseCoinsNormalized(rental.Asset)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, renter, asset)
	if err != nil {
		return nil, err
	}
	rental.State = "cancelled"
	k.SetRental(ctx, rental)
	return &types.MsgCancelRentalResponse{}, nil
}
