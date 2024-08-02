package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params:       DefaultParams(),
		Chains:       []Chain{},
		Tokens:       []Token{},
		Transactions: []Transaction{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	if err := gs.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	for _, tx := range gs.Transactions {
		if err := validateTransaction(&tx); err != nil {
			return fmt.Errorf("invalid transaction %s: %w", TransactionId(&tx), err)
		}
	}

	chains := make(map[string]struct{})
	for _, chain := range gs.Chains {
		if _, ok := chains[chain.Id]; ok {
			return fmt.Errorf("duplicate chain id: %s", chain.Id)
		} else {
			chains[chain.Id] = struct{}{}
		}

		if err := validateChain(&chain); err != nil {
			return fmt.Errorf("invalid chain %s: %w", chain.Id, err)
		}
	}

	uniqueTokens := make(map[uint64]struct{})
	for _, token := range gs.Tokens {
		if _, ok := uniqueTokens[token.Id]; ok {
			return fmt.Errorf("duplicate token id: %v", token.Id)
		} else {
			uniqueTokens[token.Id] = struct{}{}
		}

		if err := validateToken(&token); err != nil {
			return fmt.Errorf("invalid token %v: %w", token.Id, err)
		}

		for _, info := range token.Info {
			if err := validateTokenInfo(&info, nil); err != nil {
				return fmt.Errorf("invalid token info for token %v: %w", token.Id, err)
			}
		}
	}
	return nil
}
