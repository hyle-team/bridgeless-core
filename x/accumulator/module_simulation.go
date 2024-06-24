package accumulator

//
////
////import (
////	"math/rand"
////
////	//"github.com/evmos/ethermint/testutil/sample"
////	bruhaccumulatorsimulation "github.com/evmos/ethermint/x/accumulator/simulation"
////	"github.com/evmos/ethermint/x/accumulator/types"
////	"github.com/cosmos/cosmos-sdk/baseapp"
////	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
////	sdk "github.com/cosmos/cosmos-sdk/types"
////	"github.com/cosmos/cosmos-sdk/types/module"
////	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
////	"github.com/cosmos/cosmos-sdk/x/simulation"
////)
////
////// avoid unused import issue
////var (
////	//_ = sample.AccAddress
////	_ = bruhaccumulatorsimulation.FindAccount
////	_ = simappparams.StakePerAccount
////	_ = simulation.MsgEntryKind
////	_ = baseapp.Paramspace
////)
////
////const (
////    opWeightMsgMintAllTokens = "op_weight_msg_mint_all_tokens"
//	// TODO: Determine the simulation weight value
//	defaultWeightMsgMintAllTokens int = 100
//
//	opWeightMsgMintAllTokens = "op_weight_msg_mint_all_tokens"
//	// TODO: Determine the simulation weight value
//	defaultWeightMsgMintAllTokens int = 100
//
//	opWeightMsgMintAllTokens = "op_weight_msg_mint_all_tokens"
//	// TODO: Determine the simulation weight value
//	defaultWeightMsgMintAllTokens int = 100
//
//	opWeightMsgMintAllTokens = "op_weight_msg_mint_all_tokens"
//	// TODO: Determine the simulation weight value
//	defaultWeightMsgMintAllTokens int = 100
//
//	// this line is used by starport scaffolding # simapp/module/const
////)
////
////// GenerateGenesisState creates a randomized GenState of the module
////func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
////	accs := make([]string, len(simState.Accounts))
////	for i, acc := range simState.Accounts {
////		accs[i] = acc.Address.String()
////	}
////	bruhaccumulatorGenesis := types.GenesisState{
////		Params:	types.DefaultParams(),
////		// this line is used by starport scaffolding # simapp/module/genesisState
////	}
////	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&bruhaccumulatorGenesis)
////}
////
////// ProposalContents doesn't return any content functions for governance proposals
////func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
////	return nil
////}
////
////// RandomizedParams creates randomized  param changes for the simulator
////func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
////
////	return []simtypes.ParamChange{
////	}
////}
////
////// RegisterStoreDecoder registers a decoder
////func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}
////
////// WeightedOperations returns the all the gov module operations with their respective weights.
////func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
////	operations := make([]simtypes.WeightedOperation, 0)
////
////	var weightMsgMintAllTokens int
//	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintAllTokens, &weightMsgMintAllTokens, nil,
//		func(_ *rand.Rand) {
//			weightMsgMintAllTokens = defaultWeightMsgMintAllTokens
//		},
//	)
//	operations = append(operations, simulation.NewWeightedOperation(
//		weightMsgMintAllTokens,
//		accumulatorsimulation.SimulateMsgMintAllTokens(am.accountKeeper, am.bankKeeper, am.keeper),
//	))
//
//	var weightMsgMintAllTokens int
//	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintAllTokens, &weightMsgMintAllTokens, nil,
//		func(_ *rand.Rand) {
//			weightMsgMintAllTokens = defaultWeightMsgMintAllTokens
//		},
//	)
//	operations = append(operations, simulation.NewWeightedOperation(
//		weightMsgMintAllTokens,
//		accumulatorsimulation.SimulateMsgMintAllTokens(am.accountKeeper, am.bankKeeper, am.keeper),
//	))
//
//	var weightMsgMintAllTokens int
//	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintAllTokens, &weightMsgMintAllTokens, nil,
//		func(_ *rand.Rand) {
//			weightMsgMintAllTokens = defaultWeightMsgMintAllTokens
//		},
//	)
//	operations = append(operations, simulation.NewWeightedOperation(
//		weightMsgMintAllTokens,
//		accumulatorsimulation.SimulateMsgMintAllTokens(am.accountKeeper, am.bankKeeper, am.keeper),
//	))
//
//	var weightMsgMintAllTokens int
//	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintAllTokens, &weightMsgMintAllTokens, nil,
//		func(_ *rand.Rand) {
//			weightMsgMintAllTokens = defaultWeightMsgMintAllTokens
//		},
//	)
//	operations = append(operations, simulation.NewWeightedOperation(
//		weightMsgMintAllTokens,
//		accumulatorsimulation.SimulateMsgMintAllTokens(am.accountKeeper, am.bankKeeper, am.keeper),
//	))
//
//	// this line is used by starport scaffolding # simapp/module/operation
////
////	return operations
////}
