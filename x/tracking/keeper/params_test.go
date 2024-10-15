package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/hyle-team/bridgeless-core/testutil/keeper"
	"github.com/hyle-team/bridgeless-core/x/tracking/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TrackingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
