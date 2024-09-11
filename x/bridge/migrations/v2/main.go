package v2

import (
	"cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/hyle-team/bridgeless-core/x/bridge/keeper"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func MigrateStore(ctx sdk.Context, k keeper.Keeper) error {
	ctx.Logger().Info(fmt.Sprintf("Performing v12.1.10-rc1 %s module migrations", types.ModuleName))

	// reading old structures of token info (without field "is_wrapped" into new one)
	tokens, _, err := k.GetTokensWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		return errors.Wrap(err, "failed to get tokens")
	}

	// just resetting correct structures where "is_wrapped" exists
	for _, token := range tokens {
		k.SetToken(ctx, token)
		for _, info := range token.Info {
			// adding new store
			k.SetTokenInfo(ctx, info)
			k.SetTokenPairs(ctx, info, token.Info...)
		}
	}
	return nil
}
