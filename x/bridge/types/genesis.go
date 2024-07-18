package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	if err := gs.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	chains := make(map[string]Chain)
	for _, chain := range gs.Chains {
		if _, ok := chains[chain.Id]; ok {
			return fmt.Errorf("duplicate chain id: %s", chain.Id)
		} else {
			chains[chain.Id] = chain
		}

		switch chain.Type {
		case ChainType_EVM:
			if _, set := big.NewInt(0).SetString(chain.Id, 10); !set {
				return fmt.Errorf("invalid chain id: %s", chain.Id)
			}
			if !common.IsHexAddress(chain.BridgeAddress) {
				return fmt.Errorf("invalid bridge address: %s", chain.BridgeAddress)
			}
			if !common.IsHexAddress(chain.Operator) {
				return fmt.Errorf("invalid operator address: %s", chain.Operator)
			}
		case ChainType_BITCOIN:
		case ChainType_COSMOS:
		case ChainType_OTHER:
		default:
			return fmt.Errorf("invalid chain type: %s", chain.Type)
		}
	}

	uniqueTokens := make(map[uint64]struct{})
	for _, token := range gs.Tokens {
		if _, ok := uniqueTokens[token.Id]; ok {
			return fmt.Errorf("duplicate token id: %v", token.Id)
		} else {
			uniqueTokens[token.Id] = struct{}{}
		}

		for chainStr, info := range token.Info {
			chain, ok := chains[chainStr]
			if !ok {
				return fmt.Errorf("unknown chain id: %s for token %v", chainStr, token.Id)
			}

			switch chain.Type {
			case ChainType_EVM:
				if !common.IsHexAddress(info.Address) {
					return fmt.Errorf("invalid token address: %s for token %v", info.Address, token.Id)
				}
				if info.Decimals == 0 {
					return fmt.Errorf("invalid token decimals: %v for token %v", info.Decimals, token.Id)
				}
			case ChainType_BITCOIN:
			case ChainType_COSMOS:
			case ChainType_OTHER:
			default:
				return fmt.Errorf("invalid chain type: %s", chain.Type)
			}
		}
	}

	return nil
}
