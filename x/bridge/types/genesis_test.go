package types_test

import (
	"testing"

	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "invalid genesis state",
			genState: &types.GenesisState{

				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			valid: true,
			desc:  "valid genesis state",
			genState: &types.GenesisState{
				Params: types.DefaultParams(),
				Chains: []types.Chain{
					{
						Id:            "1",
						Type:          types.ChainType_EVM,
						BridgeAddress: "0xe61174fa1b7e52132d8b365044bf95b0d90f442f",
						Operator:      "0xe61174fa1b7e52132d8b365044bf95b0d90f442f",
					},
					{
						Id:            "2",
						Type:          types.ChainType_EVM,
						BridgeAddress: "0xe61174fa1b7e52132d8b365044bf95b0d90f442f",
						Operator:      "0xe61174fa1b7e52132d8b365044bf95b0d90f442f",
					},
				},
				Tokens: []types.Token{
					{
						Id:     1,
						Name:   "test",
						Symbol: "test",
						Info: map[string]*types.TokenInfo{
							"1": {
								Decimals: 15,
								Address:  "0xe61174fa1b7e52132d8b365044bf95b0d90f442f",
							},
							"2": {
								Decimals: 35,
								Address:  "0xe61174fa1b7e52132d8b365044bf95b0d90f442f",
							},
						},
					},
				},
			},
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
