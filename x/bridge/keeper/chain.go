package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func (k Keeper) SetChain(sdkCtx sdk.Context, chain types.Chain) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreChainPrefix))
	cStore.Set(types.KeyChain(chain.Id), k.cdc.MustMarshal(&chain))
}

func (k Keeper) GetChain(sdkCtx sdk.Context, id string) (chain types.Chain, found bool) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreChainPrefix))
	bz := cStore.Get(types.KeyChain(id))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &chain)
	found = true

	return
}

func (k Keeper) GetChains(sdkCtx sdk.Context) (chains []types.Chain) {
	cStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreChainPrefix))
	iterator := cStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var chain types.Chain
		k.cdc.MustUnmarshal(iterator.Value(), &chain)
		chains = append(chains, chain)
	}

	return
}
