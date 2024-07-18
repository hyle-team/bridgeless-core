package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// GetModuleAdmin get the module admin
func (k Keeper) GetModuleAdmin(ctx sdk.Context) sdk.AccAddress {
	var admin string
	k.paramstore.Get(ctx, types.KeyPrefix(types.ParamStoreKeyModuleAdmin), &admin)

	addr, err := sdk.AccAddressFromBech32(admin)
	if err != nil {
		panic(fmt.Errorf("invalid module admin address: %s", admin))
	}

	return addr
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
