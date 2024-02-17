package keeper

import (
	"context"

	"rental/x/rental/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ApproveRental(goCtx context.Context, msg *types.MsgApproveRental) (*types.MsgApproveRentalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: Handling the message
	rental, found := k.GetRental(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if rental.State != "requested" {
		return nil, errorsmod.Wrapf(types.ErrWrongRentalState, "%v", rental.State)
	}
	owner, _ := sdk.AccAddressFromBech32(msg.Creator)
	renter, _ := sdk.AccAddressFromBech32(rental.Renter)
	amount, err := sdk.ParseCoinsNormalized(rental.Amount)

	if err != nil {
		return nil, errorsmod.Wrap(types.ErrWrongRentalState, "Cannot parse coins in rental amount")
	}
	err = k.bankKeeper.SendCoins(ctx, owner, renter, amount)
	if err != nil {
		return nil, err
	}
	rental.Owner = msg.Creator
	rental.State = "approved"
	k.SetRental(ctx, rental)
	return &types.MsgApproveRentalResponse{}, nil
}
