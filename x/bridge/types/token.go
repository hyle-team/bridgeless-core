package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	bridgeTypes "github.com/hyle-team/bridgeless-core/v12/types"
)

func validateToken(token *Token) error {
	if token == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "token is nil")
	}

	return validateTokenMetadata(&token.Metadata)
}

func validateTokenMetadata(metadata *TokenMetadata) error {
	if metadata == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidType, "metadata is nil")
	}

	if metadata.Name == "" {
		return errorsmod.Wrap(bridgeTypes.ErrInvalidMetadata, "name cannot be empty")
	}
	if metadata.Symbol == "" {
		return errorsmod.Wrap(bridgeTypes.ErrInvalidMetadata, "symbol cannot be empty")
	}
	if metadata.Uri == "" {
		return errorsmod.Wrap(bridgeTypes.ErrInvalidMetadata, "uri cannot be empty")
	}

	return nil
}

func validateTokenInfo(info *TokenInfo, chainType *ChainType) error {
	if info == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidType, "info is nil")
	}
	if info.ChainId == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidChainID, "chain id cannot be empty")

	}
	if info.Address == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "token address is empty")
	}
	if chainType == nil {
		return nil
	}

	switch *chainType {
	case ChainType_EVM:
		if !common.IsHexAddress(info.Address) {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid token address: %s", info.Address)
		}
		if info.Decimals == 0 {
			return errorsmod.Wrapf(bridgeTypes.ErrInvalidAmount, "invalid token decimals: %v", info.Decimals)
		}
	case ChainType_BITCOIN:
	case ChainType_COSMOS:
	case ChainType_ZANO:
	case ChainType_OTHER:
	default:
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid chain type: %v", *chainType)
	}

	return nil
}
