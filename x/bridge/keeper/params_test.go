package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"testing"

	testkeeper "github.com/hyle-team/bridgeless-core/testutil/keeper"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BridgeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))

	acc, _ := sdk.AccAddressFromBech32(params.ModuleAdmin)
	require.EqualValues(t, acc, k.GetModuleAdmin(ctx))
}
