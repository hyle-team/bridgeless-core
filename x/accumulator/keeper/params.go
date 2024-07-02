package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v12/x/accumulator/types"
)

// GetParams get all parameters as types.Params
func (k BaseKeeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)

	b := store.Get(types.KeyPrefix(types.ParamsKey))
	if b == nil {
		return types.DefaultParams()
	}

	k.cdc.MustUnmarshal(b, &params)
	return params
}

// SetParams set the params
func (k BaseKeeper) SetParams(ctx sdk.Context, params types.Params) {
	if err := params.Validate(); err != nil {
		panic("failed to set params: " + err.Error())
	}

	b := k.cdc.MustMarshal(&params)
	ctx.KVStore(k.storeKey).Set(types.KeyPrefix(types.ParamsKey), b)
}
