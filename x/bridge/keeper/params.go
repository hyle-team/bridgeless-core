package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	var params types.Params
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

func (k Keeper) IsParty(sender string, ctx sdk.Context) bool {
	parties := k.GetParams(ctx).Parties
	for _, party := range parties {
		if party.Address == sender {
			return true
		}
	}
	return false
}

func (k Keeper) ModuleAdmin(ctx sdk.Context) (adminAddress string) {
	k.paramstore.Get(ctx, types.ParamModuleAdminKey, &adminAddress)

	return
}

func (k Keeper) PartiesList(ctx sdk.Context) (parties []*types.Party) {
	k.paramstore.Get(ctx, types.ParamModulePartiesListKey, &parties)

	return
}

func (k *Keeper) IsAdmin(sender string, ctx sdk.Context) bool {
	return sender == k.ModuleAdmin(ctx)
}
