package types

import (
	"fmt"
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

	if tx.Receiver == "" {
		return fmt.Errorf("receiver cannot be empty")
	}

	if tx.DepositTxHash == "" {
		return fmt.Errorf("deposit tx hash cannot be empty")
	}

	return nil
}

// ValidateChainTransaction is used to validate transaction details depending on certain chain type
func ValidateChainTransaction(tx *Transaction, chainType *ChainType) error {
	if tx == nil {
		return fmt.Errorf("transaction is nil")
	}
	if chainType == nil {
		return fmt.Errorf("chain type is nil")
	}

	// Only ZANO chain can have empty depositor address
	if len(tx.DepositChainId) == 0 && *chainType != ChainType_ZANO {
		return fmt.Errorf("depositor address can`t be empty, unless it`s ZANO chain")
	}

	// Place for custom validation logic
	switch *chainType {
	case ChainType_ZANO:
	case ChainType_EVM:

	case ChainType_BITCOIN:
	case ChainType_COSMOS:
	case ChainType_OTHER:
	}

	return nil
}
