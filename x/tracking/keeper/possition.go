package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hyle-team/bridgeless-core/x/tracking/types"
	"github.com/pkg/errors"
)

func (k Keeper) SetPosition(ctx sdk.Context, positionAddress string, position types.Position) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPosition(positionAddress), k.cdc.MustMarshal(&position))
}

func (k Keeper) GetPosition(ctx sdk.Context, positionAddress string) (position types.Position, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPosition(positionAddress))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &position)
	found = true

	return
}

func (k Keeper) GetPositionsWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Position, *query.PageResponse, error) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPosition(""))
	positions := make([]types.Position, 0)
	pageRes, err := query.Paginate(store, pagination, func(key []byte, value []byte) error {
		var position types.Position
		k.cdc.MustUnmarshal(iterator.Value(), &position)

		positions = append(positions, position)
		return nil
	})

	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to paginate positions")
	}

	return positions, pageRes, err
}

func (k Keeper) DeletePosition(ctx sdk.Context, positionAddress string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyPosition(positionAddress))
}
