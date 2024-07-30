package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func (m msgServer) SubmitTransaction(goCtx context.Context, msg *types.MsgSubmitTransaction) (*types.MsgSubmitTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := m.GetTransaction(ctx, types.TransactionId(&msg.Transaction)); found {
		return nil, types.ErrTranscationAlreadySubmitted
	}
	if _, found := m.GetChain(ctx, msg.Transaction.DepositChainId); !found {
		return nil, types.ErrSourceChainNotSupported
	}

	dstChain, found := m.GetChain(ctx, msg.Transaction.WithdrawalChainId)
	if !found {
		return nil, types.ErrDestinationChainNotSupported
	}
	admin := m.GetModuleAdmin(ctx, dstChain.Type)

	if !admin.Equals(msg.GetSigners()[0]) {
		return nil, types.ErrOperationNotAllowed
	}

	m.SetTransaction(ctx, msg.Transaction)

	return &types.MsgSubmitTransactionResponse{}, nil
}
