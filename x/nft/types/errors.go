package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/nft module sentinel errors
var (
	ErrSample                  = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidAmount           = sdkerrors.Register(ModuleName, 1101, "invalid amount")
	ErrValidatorNotFound       = sdkerrors.Register(ModuleName, 1102, "validator not found")
	ErrFailedToDelegate        = sdkerrors.Register(ModuleName, 1103, "failed to delegate nft")
	ErrFailedToSendTokenAmount = sdkerrors.Register(ModuleName, 1104, "failed to send token amount")
	ErrTokenIsDelegated        = sdkerrors.Register(ModuleName, 1105, "token is delegated")
)
