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
)
