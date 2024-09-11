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
		for _, info := range token.Info {
			k.SetTokenInfo(ctx, info)
			k.SetTokenPairs(ctx, info, token.Info...)
		}
	}
	for _, tx := range genState.Transactions {
		k.SetTransaction(ctx, tx)
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	txs, _, err := k.GetPaginatedTransactions(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(fmt.Errorf("failed to export genesis transactions: %w", err))
	}

	tokens, _, err := k.GetTokensWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(fmt.Errorf("failed to export genesis tokens: %w", err))
	}

	chains, _, err := k.GetChainsWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		panic(fmt.Errorf("failed to export genesis chains: %w", err))
	}

	return &types.GenesisState{
		Params:       k.GetParams(ctx),
		Chains:       chains,
		Tokens:       tokens,
		Transactions: txs,
	}
}
