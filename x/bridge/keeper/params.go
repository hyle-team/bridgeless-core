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

func (k Keeper) GetAdmin(ctx sdk.Context, chainType types.ChainType) (sdk.AccAddress, bool) {
	var admin string

	switch chainType {
	case types.ChainType_EVM:
		k.paramstore.Get(ctx, types.KeyPrefix(types.ParamEvmAdminKey), &admin)
	default:
		return nil, false
	}

	addr, err := sdk.AccAddressFromBech32(admin)
	if err != nil {
		panic(fmt.Errorf("invalid admin address: %s", admin))
	}

	return addr, true
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
