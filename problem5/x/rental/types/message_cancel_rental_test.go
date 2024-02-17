package types

import (
	"testing"

	"rental/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCancelRental_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCancelRental
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCancelRental{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCancelRental{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
