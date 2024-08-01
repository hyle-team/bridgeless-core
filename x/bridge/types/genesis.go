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

	chains := make(map[string]Chain)
	for _, chain := range gs.Chains {
		if _, ok := chains[chain.Id]; ok {
			return fmt.Errorf("duplicate chain id: %s", chain.Id)
		} else {
			chains[chain.Id] = chain
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
	}

	for _, pair := range gs.Pairs {

		chain, ok := chains[pair.SourceChain]
		if !ok {
			return fmt.Errorf("unknown chain id: %s for token %v", pair.SourceChain, pair.TokenId)
		}

		if err := validateTokenInfo(&pair, &chain.Type); err != nil {
			return fmt.Errorf("invalid token info %v for chain %s: %w", pair.TokenId, pair.SourceChain, err)
		}
	}
	return nil
}
