package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/rental module sentinel errors
var (
	ErrInvalidSigner    = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample           = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrWrongRentalState = sdkerrors.Register(ModuleName, 1102, "wrong rental state")
	ErrDueDate          = sdkerrors.Register(ModuleName, 1103, "due date not reached")
)
