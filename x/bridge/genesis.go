package bridge

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hyle-team/bridgeless-core/x/bridge/keeper"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	for _, chain := range genState.Chains {
		k.SetChain(ctx, chain)
	}
	for _, token := range genState.Tokens {
		k.SetToken(ctx, token)
	}
	for _, tx := range genState.Transactions {
		k.SetTransaction(ctx, tx)
	}
	for _, pair := range genState.Pairs {
		k.SetTokenPair(ctx, pair.SourceChain, pair.DestinationChain, pair.Address, pair)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	txs, _, err := k.GetPaginatedTransactions(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(fmt.Errorf("failed to export genesis transactions: %w", err))
	}

	return &types.GenesisState{
		Params:       k.GetParams(ctx),
		Chains:       k.GetAllChains(ctx),
		Tokens:       k.GetAllTokens(ctx),
		Pairs:        k.GetTokenPairs(ctx),
		Transactions: txs,
	}
}
