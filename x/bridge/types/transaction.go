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

	if _, set := big.NewInt(0).SetString(tx.DepositChainId, 10); !set {
		return fmt.Errorf("invalid deposit chain id: %s", tx.DepositChainId)
	}

	if _, set := big.NewInt(0).SetString(tx.WithdrawalChainId, 10); !set {
		return fmt.Errorf("invalid withdrawal chain id: %s", tx.WithdrawalChainId)
	}

	if _, set := big.NewInt(0).SetString(tx.Amount, 10); !set {
		return fmt.Errorf("invalid amount: %s", tx.Amount)
	}

	if tx.Depositor == "" {
		return fmt.Errorf("depositor cannot be empty")
	}

	if tx.Receiver == "" {
		return fmt.Errorf("receiver cannot be empty")
	}

	if tx.DepositTxHash == "" {
		return fmt.Errorf("deposit tx hash cannot be empty")
	}

	if tx.WithdrawalTxHash == "" {
		return fmt.Errorf("withdrawal tx hash cannot be empty")
	}

	return nil
}
