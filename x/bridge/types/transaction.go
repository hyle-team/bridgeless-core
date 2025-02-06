package types

import (
	errorsmod "cosmossdk.io/errors"
	"errors"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math/big"
)

func TransactionId(tx *Transaction) string {
	return fmt.Sprintf("%s/%v/%s", tx.DepositTxHash, tx.DepositTxIndex, tx.DepositChainId)
}

func validateTransaction(tx *Transaction) error {
	if tx == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "transaction is nil")
	}

	if _, set := big.NewInt(0).SetString(tx.DepositAmount, 10); !set {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid deposit amount: %s", tx.DepositAmount)
	}

	if _, set := big.NewInt(0).SetString(tx.WithdrawalAmount, 10); !set {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid withdrawal amount: %s", tx.WithdrawalAmount)
	}

	if len(tx.DepositChainId) == 0 || len(tx.WithdrawalChainId) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "deposit chain id and withdrawal id cannot be empty")
	}
	if tx.WithdrawalChainId == tx.DepositChainId {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "withdrawal chain id cannot be the same as deposit chain id ")
	}

	if tx.Receiver == "" {
		return errorsmod.Wrap(errors.New("invalid receiver"), "receiver cannot be empty")
	}

	if tx.DepositTxHash == "" {
		return errorsmod.Wrap(errors.New("invalid tx hash"), "deposit tx hash cannot be empty")
	}

	return nil
}

// ValidateChainTransaction is used to validate transaction details depending on certain chain type
func ValidateChainTransaction(tx *Transaction, chainType ChainType) error {
	if tx == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "transaction is nil")
	}
	if len(chainType.String()) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "chain type is empty")
	}

	// Only ZANO chain can have empty depositor address
	if len(tx.Depositor) == 0 && chainType != ChainType_ZANO {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "depositor address can`t be empty, unless it`s ZANO chain")
	}

	// Place for custom validation logic
	switch chainType {
	case ChainType_EVM:
	case ChainType_ZANO:
	case ChainType_BITCOIN:
	case ChainType_COSMOS:
	case ChainType_OTHER:
	}

	return nil
}
