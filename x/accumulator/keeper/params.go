package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v12/x/accumulator/types"
)

// GetParams get all parameters as types.Params
func (k BaseKeeper) GetParams(ctx sdk.Context) (params types.Params) {
	fmt.Println("getting params")
	store := ctx.KVStore(k.storeKey)
	fmt.Println("store ", store)

	b := store.Get(types.KeyPrefix(types.ParamsKey))
	fmt.Println("b ", b)

	if b == nil {
		fmt.Println("b is nil")
		return types.DefaultParams()
	}

	fmt.Println("k.cdc", k.cdc)
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
