package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetToken(sdkCtx sdk.Context, token types.Token) {
	tStore := prefix.NewStore(sdkCtx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPrefix))
	tStore.Set(types.KeyToken(token.Id), k.cdc.MustMarshal(&token))
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
}

func (k Keeper) GetAllTokens(sdkCtx sdk.Context) (tokens []types.Token) {
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

func (k Keeper) GetTokensWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Token, *query.PageResponse, error) {
	tStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoreTokenPrefix))
	var chains []types.Token

	pageRes, err := query.Paginate(tStore, pagination, func(key []byte, value []byte) error {
		var chain types.Token

		k.cdc.MustUnmarshal(value, &chain)

		chains = append(chains, chain)
		return nil
	})

	if err != nil {
		return nil, pageRes, status.Error(codes.Internal, err.Error())
	}

	return chains, pageRes, nil
}
