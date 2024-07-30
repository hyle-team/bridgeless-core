package keeper

import (
	"cosmossdk.io/errors"
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func (k Keeper) SetToken(sdkCtx sdk.Context, token types.Token) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPrefix))
	tStore.Set(types.KeyToken(token.Id), k.cdc.MustMarshal(&token))

	for chain, info := range token.Info {
		if info != nil {
			k.SetTokenPairs(sdkCtx, chain, info.Address, token.Info)
		}
	}
}

func (k Keeper) GetToken(sdkCtx sdk.Context, id uint64) (token types.Token, found bool) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPrefix))
	bz := tStore.Get(types.KeyToken(id))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &token)
	found = true

	return
}

func (k Keeper) RemoveToken(sdkCtx sdk.Context, id uint64) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPrefix))
	tStore.Delete(types.KeyToken(id))

	// TODO: rm token pairs
}

func (k Keeper) GetTokens(sdkCtx sdk.Context) (tokens []types.Token) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPrefix))
	iterator := tStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var token types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &token)
		tokens = append(tokens, token)
	}

	return
}

func (k Keeper) SetTokenPairs(sdkCtx sdk.Context, srcChain, srcAddress string, info map[string]*types.TokenInfo) {
	raw, err := json.Marshal(info)
	if err != nil {
		panic(errors.Wrap(err, "failed to marshal token info"))
	}

	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	pStore.Set(types.KeyTokenPair(srcChain, srcAddress), raw)
}

func (k Keeper) GetTokenPairs(sdkCtx sdk.Context, srcChain, srcAddress string) (info map[string]*types.TokenInfo, found bool) {
	pStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPairsPrefix))
	bz := pStore.Get(types.KeyTokenPair(srcChain, srcAddress))
	if bz == nil {
		return
	}

	if err := json.Unmarshal(bz, &info); err != nil {
		panic(errors.Wrap(err, "failed to unmarshal token info"))
	}

	found = true

	return
}

func (k Keeper) GetTokenPair(sdk sdk.Context, srcChain, srcAddress, dstChain string) (info *types.TokenInfo, found bool) {
	pairs, found := k.GetTokenPairs(sdk, srcChain, srcAddress)
	if !found {
		return
	}

	info, found = pairs[dstChain]

	return
}
