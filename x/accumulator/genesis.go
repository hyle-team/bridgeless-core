package accumulator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v18/x/accumulator/keeper"
	"github.com/evmos/evmos/v18/x/accumulator/types"
	"time"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	for key, m := range genState.Params.ModulesInfo {
		err := k.MintTokens(ctx, m.Amount, key)
		if err != nil {
			k.Logger(ctx).Error(err.Error())
		}

		if m.Vesting != nil {
			vesting := keeper.NewBaseAdminVesting(k, m.Address, time.Time{}, m.Vesting)
			err := vesting.UnlockImmediately(m.Amount - m.Vesting.TotalLockedAmount)
			if err != nil {
				k.Logger(ctx).Error(err.Error())
				return
			}
		}
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
