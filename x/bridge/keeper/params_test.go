package keeper_test

import (
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
}
