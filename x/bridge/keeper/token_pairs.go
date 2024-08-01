package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func (k Keeper) SetTokenPair(sdkCtx sdk.Context, srcChain, dstChain, srcAddress string, pair *types.TokenInfo) {
	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.KeyPrefix(fmt.Sprintf("%s-%s", srcChain, srcAddress)))

	srcBranchStore.Set(types.KeyTokenPair(dstChain), k.cdc.MustMarshal(pair))
}

func (k Keeper) GetTokenPair(sdkCtx sdk.Context, srcChain, dstChain, srcAddress string) (info *types.TokenInfo, found bool) {
	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.KeyPrefix(fmt.Sprintf("%s-%s", srcChain, srcAddress)))

	bz := srcBranchStore.Get(types.KeyTokenPair(dstChain))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, info)
	return info, true
}

func (k Keeper) GetTokenPairs(ctx sdk.Context, srcChain, srcAddress string) (list []*types.TokenInfo) {
	pStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.KeyPrefix(fmt.Sprintf("%s-%s", srcChain, srcAddress)))
	iterator := sdk.KVStorePrefixIterator(srcBranchStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val *types.TokenInfo
		k.cdc.MustUnmarshal(iterator.Value(), val)
		list = append(list, val)
	}

	return
}

func (k Keeper) RemoveTokenPairs(sdkCtx sdk.Context, srcChain, srcAddress, dstChain string) {
	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.KeyPrefix(fmt.Sprintf("%s-%s", srcChain, srcAddress)))

	srcBranchStore.Delete(types.KeyTokenPair(dstChain))
}
