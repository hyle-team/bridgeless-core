package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func validateChain(chain *Chain) error {
	if chain == nil {
		return fmt.Errorf("chain is nil")
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
	case ChainType_ZANO:
	default:
		return fmt.Errorf("invalid chain type: %s", chain.Type)
	}

	return nil
}
