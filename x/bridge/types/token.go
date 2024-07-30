package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func validateToken(token *Token) error {
	if token == nil {
		return fmt.Errorf("token is nil")
	}

	if token.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if token.Symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}

	return nil
}

func validateTokenInfo(info *TokenInfo, chainType *ChainType) error {
	if info == nil {
		return fmt.Errorf("info is nil")
	}
	if chainType == nil {
		return nil
	}

	switch *chainType {
	case ChainType_EVM:
		if !common.IsHexAddress(info.Address) {
			return fmt.Errorf("invalid token address: %s", info.Address)
		}
		if info.Decimals == 0 {
			return fmt.Errorf("invalid token decimals: %v", info.Decimals)
		}
	case ChainType_BITCOIN:
	case ChainType_COSMOS:
	case ChainType_OTHER:
	default:
		return fmt.Errorf("invalid chain type: %v", *chainType)
	}

	return nil
}
