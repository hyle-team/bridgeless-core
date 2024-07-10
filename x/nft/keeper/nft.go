package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/nft/types"
)

// SetNFT set a specific deposit in the store from its index
func (k Keeper) SetNFT(ctx sdk.Context, v types.NFT) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTKeyPrefix))
	b := k.cdc.MustMarshal(&v)
	store.Set(types.NFTKey(v.Address), b)
}

// GetNFT returns a NFT from its index
func (k Keeper) GetNFT(
	ctx sdk.Context,
	address string,
) (val types.NFT, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTKeyPrefix))

	b := store.Get(types.NFTKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFT removes a NFT from the store
func (k Keeper) RemoveNFT(
	ctx sdk.Context,
	address string,
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTKeyPrefix))
	store.Delete(types.NFTKey(
		address,
	))
}

// GetAllNFT returns all NFT
func (k Keeper) GetAllNFT(ctx sdk.Context) (list []types.NFT) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.NFT
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
