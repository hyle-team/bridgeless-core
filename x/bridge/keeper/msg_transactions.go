package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func (m msgServer) SubmitTransaction(goCtx context.Context, msg *types.MsgSubmitTransaction) (*types.MsgSubmitTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Submitter != m.GetParams(ctx).ModuleAdmin {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	if _, found := m.GetTransaction(ctx, types.TransactionId(&msg.Transaction)); found {
		return nil, types.ErrTranscationAlreadySubmitted
	}
	if _, found := m.GetChain(ctx, msg.Transaction.DepositChainId); !found {
		return nil, types.ErrSourceChainNotSupported
	}
	if _, found := m.GetChain(ctx, msg.Transaction.WithdrawalChainId); !found {
		return nil, types.ErrDestinationChainNotSupported
	}

	m.SetTransaction(ctx, msg.Transaction)

	return &types.MsgSubmitTransactionResponse{}, nil
}
