package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func (k Keeper) GetDstToken(sdkCtx sdk.Context, srcAddr, srcChain, dscChain string) (info types.TokenInfo, found bool) {
	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.TokenPairPrexif(srcChain, srcAddr))

	bz := srcBranchStore.Get(types.KeyTokenPair(dscChain))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &info)
	found = true

	return
}

func (k Keeper) setTokenPairs(sdkCtx sdk.Context, current types.TokenInfo, pairs []types.TokenInfo) {
	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.TokenPairPrexif(current.ChainId, current.Address))

	for _, pair := range pairs {
		if pair.ChainId == current.ChainId {
			continue
		}

		srcBranchStore.Set(types.KeyTokenPair(pair.ChainId), k.cdc.MustMarshal(&pair))
	}
}

func (k Keeper) removeTokenPairs(sdkCtx sdk.Context, current types.TokenInfo, pairs []types.TokenInfo) {
	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.TokenPairPrexif(current.ChainId, current.Address))

	for _, pair := range pairs {
		if pair.ChainId == current.ChainId {
			continue
		}

		srcBranchStore.Delete(types.KeyTokenPair(pair.ChainId))
	}
}
