package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k BaseKeeper) EndBlocker(ctx sdk.Context) {
	module := k.GetParams(ctx).ModulesInfo["accumulator"]
	vesting := NewBaseAdminVesting(k, module.Address, k.lastVestingTime, module.Vesting)
	lastTime, err := vesting.Unlock()
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return
	}

	k.SetLastVestingTime(*lastTime)
}
