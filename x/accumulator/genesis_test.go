package accumulator_test

//
//import (
//	"testing"
//
//	//keepertest "github.com/evmos/evmos/v18/testutil/keeper"
//	//"github.com/evmos/evmos/v18/testutil/nullify"
//	"github.com/evmos/evmos/v18/x/accumulator"
//	"github.com/evmos/evmos/v18/x/accumulator/types"
//	"github.com/stretchr/testify/require"
//)
//
//func TestGenesis(t *testing.T) {
//	genesisState := types.GenesisState{
//		Params: types.DefaultParams(),
//
//		// this line is used by starport scaffolding # genesis/test/state
//	}
//
//	k, ctx := keepertest.BruhaccumulatorKeeper(t)
//	accumulator.InitGenesis(ctx, *k, genesisState)
//	got := accumulator.ExportGenesis(ctx, *k)
//	require.NotNil(t, got)
//
//	nullify.Fill(&genesisState)
//	nullify.Fill(got)
//
//	// this line is used by starport scaffolding # genesis/test/assert
//}
