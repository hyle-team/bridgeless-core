package keeper_test

import (
	"testing"

	testkeeper "github.com/hyle-team/bridgeless-core/v12/testutil/keeper"
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TrackingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
