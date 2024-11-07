package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (m msgServer) SubmitTransactions(goCtx context.Context, msg *types.MsgSubmitTransactions) (*types.MsgSubmitTransactionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Submitter != m.GetParams(ctx).ModuleAdmin {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	for _, tx := range msg.Transactions {
		if _, found := m.GetTransaction(ctx, types.TransactionId(&tx)); found {
			return nil, types.ErrTranscationAlreadySubmitted
		}
		if _, found := m.GetChain(ctx, tx.DepositChainId); !found {
			return nil, types.ErrSourceChainNotSupported
		}
		if _, found := m.GetChain(ctx, tx.WithdrawalChainId); !found {
			return nil, types.ErrDestinationChainNotSupported
		}

		m.SetTransaction(ctx, tx)
	}

	return &types.MsgSubmitTransactionsResponse{}, nil
}
