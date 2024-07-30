package bridge_test

import (
	keepertest "github.com/hyle-team/bridgeless-core/testutil/keeper"
	"github.com/hyle-team/bridgeless-core/x/bridge"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenesis_HappyPath(t *testing.T) {
	genesisState := types.DefaultGenesis()
	genesisState.Chains = []types.Chain{
		{
			Id:            "123",
			Type:          types.ChainType_EVM,
			BridgeAddress: "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627",
			Operator:      "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627",
		},
		{
			Id:            "567738",
			Type:          types.ChainType_EVM,
			BridgeAddress: "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627",
			Operator:      "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627",
		},
		{
			Id:   "567739",
			Type: types.ChainType_OTHER,
		},
	}
	genesisState.Tokens = []types.Token{
		{
			Name:   "USDToken",
			Symbol: "USDT",
			Info: map[string]*types.TokenInfo{
				"123": {
					Address:  "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627",
					Decimals: 18,
				},
				"567738": {
					Address:  "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627",
					Decimals: 8,
				},
				"567739": {},
			},
		},
		{
			Id:     1,
			Name:   "USDToken1",
			Symbol: "USDT",
			Info: map[string]*types.TokenInfo{
				"123": {
					Address:  "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b628",
					Decimals: 18,
				},
				"567738": {
					Address:  "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b628",
					Decimals: 8,
				},
				"567739": {},
			},
		},
	}
	genesisState.Transactions = []types.Transaction{
		{
			DepositChainId:    "8002",
			DepositTxHash:     "0x764052fe388caefececd77a0b4bc64fda6738f915611cac0dac3bc6ea1646908",
			DepositTxIndex:    1,
			Amount:            "1000000000000000000",
			Depositor:         "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627",
			Receiver:          "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627",
			DepositBlock:      123,
			WithdrawalTxHash:  "0x764052fe388caefececd77a0b4bc64fda6738f915611cac0dac3bc6ea1646908",
			WithdrawalChainId: "8002",
		},
	}

	require.Nil(t, genesisState.Validate())

	k, ctx := keepertest.BridgeKeeper(t)
	bridge.InitGenesis(ctx, *k, *genesisState)

	chain, found := k.GetChain(ctx, "123")
	require.True(t, found)
	require.Equal(t, "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627", chain.BridgeAddress)

	chain, found = k.GetChain(ctx, "5677eqwe")
	require.False(t, found)

	token, found := k.GetToken(ctx, 1)
	require.True(t, found)
	require.Equal(t, "USDToken1", token.Name)

	token, found = k.GetToken(ctx, 2)
	require.False(t, found)

	tokenPairs, found := k.GetTokenPairs(ctx, "123", "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627")
	require.True(t, found)
	require.Equal(t, 3, len(tokenPairs))

	tokenPair, found := k.GetTokenPair(ctx, "123", "0x9c9b83Ed9dd4cF8A385b6e318Fb97Cdfc320b627", "567738")
	require.True(t, found)
	require.Equal(t, uint64(8), tokenPair.Decimals)

	got := bridge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)
	require.EqualValues(t, genesisState, got)

	//this line is used by starport scaffolding # genesis/test/assert
}
