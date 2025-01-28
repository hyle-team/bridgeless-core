package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"math/big"
	"strconv"
)

func (k Keeper) SetTransaction(sdkCtx sdk.Context, transaction types.Transaction) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionPrefix))
	tStore.Set(types.KeyTransaction(types.TransactionId(&transaction)), k.cdc.MustMarshal(&transaction))

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

func (k Keeper) GetTransaction(sdkCtx sdk.Context, id string) (types.Transaction, bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionPrefix))

	var transaction types.Transaction
	bz := tStore.Get(types.KeyTransaction(id))
	if bz == nil {
		return transaction, false
	}

	k.cdc.MustUnmarshal(bz, &transaction)
	return transaction, true
}

func (k Keeper) GetPaginatedTransactions(
	sdkCtx sdk.Context, pagination *query.PageRequest,
) (
	[]types.Transaction, *query.PageResponse, error,
) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionPrefix))

	var transactions []types.Transaction
	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var transaction types.Transaction
		if err := k.cdc.Unmarshal(value, &transaction); err != nil {
			return err
		}
		transactions = append(transactions, transaction)
		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return transactions, pageRes, nil
}
