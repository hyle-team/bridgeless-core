package nft_test

import (
	"testing"

	keepertest "github.com/evmos/evmos/v12/testutil/keeper"
	"github.com/evmos/evmos/v12/testutil/nullify"
	"github.com/evmos/evmos/v12/x/nft"
	"github.com/evmos/evmos/v12/x/nft/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NftKeeper(t)
	nft.InitGenesis(ctx, *k, genesisState)
	got := nft.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
