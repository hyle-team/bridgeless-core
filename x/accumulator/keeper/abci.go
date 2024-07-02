package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v12/x/accumulator/types"
)

func (k BaseKeeper) EndBlocker(ctx sdk.Context) {
	module := k.GetModuleInfo(ctx, types.ModuleName)
	if module == nil {
		k.Logger(ctx).Error("module info not found")
		return
	}
	vesting := NewBaseAdminVesting(k, module.Address, k.lastVestingTime, module.Vesting)
	lastTime, err := vesting.Unlock()
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return
	}

	k.SetLastVestingTime(*lastTime)
}
