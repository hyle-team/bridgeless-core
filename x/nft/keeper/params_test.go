package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/evmos/evmos/v12/testutil/keeper"
	"github.com/evmos/evmos/v12/x/nft/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NftKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
