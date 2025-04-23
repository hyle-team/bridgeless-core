package keeper

import (
	"bytes"
	"encoding/gob"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (k Keeper) SetTransactionSubmitters(sdkCtx sdk.Context, transaction *types.Transaction, submitters []string) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionSubmissionsPrefix))
	buf := new(bytes.Buffer)
	gob.NewEncoder(buf).Encode(submitters)
	tStore.Set(k.TxHash(transaction).Bytes(), buf.Bytes())

}

func (k Keeper) GetTransactionSubmitters(sdkCtx sdk.Context, hash common.Hash) (submitters []string) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.Prefix(types.StoreTransactionSubmissionsPrefix))

	bz := tStore.Get(hash.Bytes())
	if bz == nil {
		return submitters
	}

	gob.NewDecoder(bytes.NewBuffer(bz)).Decode(&submitters)

	return submitters
}

func (k Keeper) TxHash(tx *types.Transaction) common.Hash {
	return crypto.Keccak256Hash(k.cdc.MustMarshal(tx))
}
