package types

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

func validateToken(token *Token) error {
	if token == nil {
		return fmt.Errorf("token is nil")
	}

	return validateTokenMetadata(&token.Metadata)
}

func validateTokenMetadata(metadata *TokenMetadata) error {
	if metadata == nil {
		return fmt.Errorf("metadata is nil")
	}

	if metadata.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if metadata.Symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}
	if metadata.Uri == "" {
		return fmt.Errorf("uri cannot be empty")
	}

	return nil
}

func validateTokenInfo(info *TokenInfo, chainType *ChainType) error {
	if info == nil {
		return fmt.Errorf("info is nil")
	}
	if info.ChainId == "" {
		return fmt.Errorf("chain id cannot be empty")

	}
	if info.Address == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token address is empty")
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
	case ChainType_ZANO:
	case ChainType_OTHER:
	default:
		return fmt.Errorf("invalid chain type: %v", *chainType)
	}

	return nil
}
