package types

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRequestRental{}

func NewMsgRequestRental(creator string, amount string, fee string, asset string, dueDate string) *MsgRequestRental {
	return &MsgRequestRental{
		Creator: creator,
		Amount:  amount,
		Fee:     fee,
		Asset:   asset,
		DueDate: dueDate,
	}
}

func (msg *MsgRequestRental) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	amount, _ := sdk.ParseCoinsNormalized(msg.Amount)
	if !amount.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount is not a valid Coins object")
	}
	if amount.Empty() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount is empty")
	}
	fee, _ := sdk.ParseCoinsNormalized(msg.Fee)
	if !fee.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "fee is not a valid Coins object")
	}
	dueDate, err := strconv.ParseInt(msg.DueDate, 10, 64)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dute date is not an integer")
	}
	if dueDate <= 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dute date should be a positive integer")
	}
	asset, _ := sdk.ParseCoinsNormalized(msg.Asset)
	if !asset.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "asset is not a valid Coins object")
	}
	if asset.Empty() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "asset is empty")
	}
	return nil
}
