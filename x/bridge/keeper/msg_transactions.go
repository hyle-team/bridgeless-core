package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"math/big"
	"strconv"
)

func (m msgServer) SubmitTransactions(goCtx context.Context, msg *types.MsgSubmitTransactions) (*types.MsgSubmitTransactionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !m.IsParty(msg.Submitter, ctx) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "submitter isn`t an authorized party")
	}

	for _, tx := range msg.Transactions {
		if _, found := m.GetTransaction(ctx, types.TransactionId(&tx)); found {
			return nil, types.ErrTranscationAlreadySubmitted
		}
		chain, found := m.GetChain(ctx, tx.DepositChainId)
		if !found {
			return nil, types.ErrSourceChainNotSupported
		}
		if _, found = m.GetChain(ctx, tx.WithdrawalChainId); !found {
			return nil, types.ErrDestinationChainNotSupported
		}

		// Custom validation of transaction for certain chain type
		err := types.ValidateChainTransaction(&tx, chain.Type)
		if err != nil {
			return nil, sdkerrors.Wrap(types.InvalidTransaction, err.Error())
		}

		m.SetTransaction(ctx, tx)

		// emit submit deposit event to notify users about new submitted deposit
		emitSubmitEvent(ctx, tx)

	}

	return &types.MsgSubmitTransactionsResponse{}, nil
}

func emitSubmitEvent(sdkCtx sdk.Context, transaction types.Transaction) {
	sdkCtx.EventManager().EmitEvent(sdk.NewEvent(types.EventType_DEPOSIT_SUBMITTED.String(),
		sdk.NewAttribute(types.AttributeKeyDepositTxHash, transaction.DepositTxHash),
		sdk.NewAttribute(types.AttributeKeyDepositNonce, big.NewInt(int64(transaction.DepositTxIndex)).String()),
		sdk.NewAttribute(types.AttributeKeyDepositChainId, transaction.DepositChainId),
		sdk.NewAttribute(types.AttributeKeyDepositAmount, transaction.DepositAmount),
		sdk.NewAttribute(types.AttributeKeyDepositBlock, big.NewInt(int64(transaction.DepositBlock)).String()),
		sdk.NewAttribute(types.AttributeKeyDepositToken, transaction.DepositToken),
		sdk.NewAttribute(types.AttributeKeyWithdrawalAmount, transaction.WithdrawalAmount),
		sdk.NewAttribute(types.AttributeKeyDepositor, transaction.Depositor),
		sdk.NewAttribute(types.AttributeKeyReceiver, transaction.Receiver),
		sdk.NewAttribute(types.AttributeKeyWithdrawalChainID, transaction.WithdrawalChainId),
		sdk.NewAttribute(types.AttributeKeyWithdrawalTxHash, transaction.WithdrawalTxHash),
		sdk.NewAttribute(types.AttributeKeyWithdrawalToken, transaction.WithdrawalToken),
		sdk.NewAttribute(types.AttributeKeySignature, transaction.Signature),
		sdk.NewAttribute(types.AttributeKeyIsWrapped, strconv.FormatBool(transaction.IsWrapped))))
}
