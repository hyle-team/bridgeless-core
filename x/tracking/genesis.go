package tracking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/keeper"
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	// import positions
	for _, position := range genState.Positions {
		k.SetPosition(ctx, position.Address, position)
	}

}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// export positions
	positions, _, err := k.GetPositionsWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(sdkerrors.Wrap(err, "failed to export positions"))
	}
	genesis.Positions = positions

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
