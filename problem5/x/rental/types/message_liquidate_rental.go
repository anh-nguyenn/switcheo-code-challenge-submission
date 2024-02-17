package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLiquidateRental{}

func NewMsgLiquidateRental(creator string, id uint64) *MsgLiquidateRental {
	return &MsgLiquidateRental{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgLiquidateRental) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
