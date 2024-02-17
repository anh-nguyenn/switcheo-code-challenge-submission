package keeper

import (
	"context"
	"strconv"

	"rental/x/rental/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) LiquidateRental(goCtx context.Context, msg *types.MsgLiquidateRental) (*types.MsgLiquidateRentalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	rental, found := k.GetRental(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if rental.Owner != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot asset: not the owner")
	}
	if rental.State != "approved" {
		return nil, errorsmod.Wrapf(types.ErrWrongRentalState, "%v", rental.State)
	}
	owner, _ := sdk.AccAddressFromBech32(rental.Owner)
	asset, _ := sdk.ParseCoinsNormalized(rental.Asset)
	dueDate, err := strconv.ParseInt(rental.DueDate, 10, 64)
	if err != nil {
		panic(err)
	}
	if ctx.BlockHeight() < dueDate {
		return nil, errorsmod.Wrap(types.ErrDueDate, "Cannot liquidate before deadline")
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, asset)
	if err != nil {
		return nil, err
	}
	rental.State = "asseted"
	k.SetRental(ctx, rental)
	return &types.MsgLiquidateRentalResponse{}, nil
}
