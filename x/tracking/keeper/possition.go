package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SetPosition(ctx sdk.Context, positionAddress string, position types.Position) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKeyPrefix))
	b := k.cdc.MustMarshal(&position)
	store.Set(types.KeyPosition(positionAddress), b)
}

func (k Keeper) GetPosition(ctx sdk.Context, positionAddress string) (position types.Position, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKeyPrefix))
	bz := store.Get(types.KeyPosition(positionAddress))
	if bz == nil {
		return
	}

	k.cdc.MustUnmarshal(bz, &position)
	found = true

	return
}

func (k Keeper) GetPositionsWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Position, *query.PageResponse, error) {
	positions := make([]types.Position, 0)
	store := ctx.KVStore(k.storeKey)
	pStore := prefix.NewStore(store, types.KeyPrefix(types.PositionKeyPrefix))

	pageRes, err := query.Paginate(pStore, pagination, func(key []byte, value []byte) error {
		var position types.Position
		k.cdc.MustUnmarshal(value, &position)

		positions = append(positions, position)
		return nil
	})

	if err != nil {
		return nil, nil, status.Error(codes.Internal, err.Error())
	}

	return positions, pageRes, err
}

func (k Keeper) DeletePosition(ctx sdk.Context, positionAddress string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PositionKeyPrefix))
	store.Delete(types.KeyPosition(positionAddress))
}
