package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/bridge module sentinel errors
var (
	ErrSample                       = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrSourceChainNotSupported      = sdkerrors.Register(ModuleName, 1101, "source chain not supported")
	ErrDestinationChainNotSupported = sdkerrors.Register(ModuleName, 1102, "destination chain not supported")
	ErrTranscationAlreadySubmitted  = sdkerrors.Register(ModuleName, 1103, "transaction already submitted")
	ErrOperationNotAllowed          = sdkerrors.Register(ModuleName, 1104, "operation not allowed")
	ErrTokenNotFound                = sdkerrors.Register(ModuleName, 1105, "token not found")
	ErrTokenPairNotFound            = sdkerrors.Register(ModuleName, 1106, "token pair not found")
	ErrTokenInfoNotFound            = sdkerrors.Register(ModuleName, 1109, "token info not found")
	ErrChainNotFound                = sdkerrors.Register(ModuleName, 1107, "chain not found")
	ErrPermissionDenied             = sdkerrors.Register(ModuleName, 1108, "permission denied")
	InvalidTransaction              = sdkerrors.Register(ModuleName, 1110, "invalid transaction")
)
