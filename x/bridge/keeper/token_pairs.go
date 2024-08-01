package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func (k Keeper) SetTokenInfo(sdkCtx sdk.Context, srcChain, dstChain, srcAddress string, pair types.TokenInfo) error {
	token, found := k.GetToken(sdkCtx, pair.TokenId)
	if !found {
		return types.ErrTokenNotFound
	}

	token.Info = append(token.Info, pair)
	k.SetToken(sdkCtx, token)

	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.KeyPrefix(fmt.Sprintf("%s-%s", srcChain, srcAddress)))

	srcBranchStore.Set(types.KeyTokenPair(dstChain), k.cdc.MustMarshal(&pair))

	return nil
}

func (k Keeper) GetTokenInfo(sdkCtx sdk.Context, srcChain, dstChain, srcAddress string) (info types.TokenInfo, found bool) {
	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.KeyPrefix(fmt.Sprintf("%s-%s", srcChain, srcAddress)))

	bz := srcBranchStore.Get(types.KeyTokenPair(dstChain))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &info)
	return info, true
}

func (k Keeper) GetBaseTokenInfo(sdkCtx sdk.Context, tokenId uint64) types.Token {
	token, found := k.GetToken(sdkCtx, tokenId)
	if !found {
		return token
	}

	token.Info = make([]types.TokenInfo, 0)
	return token
}

func (k Keeper) GetTokenPairsByTokenId(ctx sdk.Context) (list []types.TokenInfo) {
	pStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	iterator := sdk.KVStorePrefixIterator(pStore, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TokenInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) RemoveTokenPair(sdkCtx sdk.Context, srcChain, srcAddress, dstChain string) error {
	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	srcBranchStore := prefix.NewStore(pStore, types.KeyPrefix(fmt.Sprintf("%s-%s", srcChain, srcAddress)))

	pair, found := k.GetTokenInfo(sdkCtx, srcChain, dstChain, srcAddress)
	if !found {
		return types.ErrTokenPairsNotFound
	}

	token, found := k.GetToken(sdkCtx, pair.TokenId)
	if !found {
		return types.ErrTokenNotFound
	}

	for id, tokenInfo := range token.Info {
		if tokenInfo.Address == pair.Address && tokenInfo.DestinationChain == dstChain {
			token.Info = append(token.Info[:id], token.Info[id+1:]...)
		}
	}

	k.SetToken(sdkCtx, token)

	srcBranchStore.Delete(types.KeyTokenPair(dstChain))

	return nil
}
