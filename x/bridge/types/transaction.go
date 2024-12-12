package types

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"math/big"
)

func TransactionId(tx *Transaction) string {
	return fmt.Sprintf("%s/%v/%s", tx.DepositTxHash, tx.DepositTxIndex, tx.DepositChainId)
}

func validateTransaction(tx *Transaction) error {
	if tx == nil {
		return fmt.Errorf("transaction is nil")
	}

	if _, set := big.NewInt(0).SetString(tx.Amount, 10); !set {
		return fmt.Errorf("invalid amount: %s", tx.Amount)
	}

	if len(tx.DepositChainId) == 0 || len(tx.WithdrawalChainId) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "deposit chain id and withdrawal id cannot be empty")
	}
	if tx.WithdrawalChainId == tx.DepositChainId {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "withdrawal chain id cannot be the same as deposit chain id ")
	}

	if tx.Receiver == "" {
		return fmt.Errorf("receiver cannot be empty")
	}

	if tx.DepositTxHash == "" {
		return fmt.Errorf("deposit tx hash cannot be empty")
	}

	return nil
}

// ValidateChainTransaction is used to validate transaction details depending on certain chain type
func ValidateChainTransaction(tx *Transaction, chainType ChainType) error {
	if tx == nil {
		return fmt.Errorf("transaction is nil")
	}
	if len(chainType.String()) == 0 {
		return fmt.Errorf("chain type is empty")
	}

	// Only ZANO chain can have empty depositor address
	if len(tx.Depositor) == 0 && chainType != ChainType_ZANO {
		return fmt.Errorf("depositor address can`t be empty, unless it`s ZANO chain")
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
