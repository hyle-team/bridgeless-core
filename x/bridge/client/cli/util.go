package cli

import (
	"cosmossdk.io/errors"
	"encoding/json"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"os"
)

func parseSubmitTx(path string) ([]types.Transaction, error) {
	var txs []types.Transaction
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "error reading file")
	}
	if err = json.Unmarshal(contents, &txs); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal transaction")
	}
	return txs, nil
}

func parseInsertToken(path string) (types.Token, error) {
	var token types.Token
	contents, err := os.ReadFile(path)
	if err != nil {
		return token, errors.Wrap(err, "error reading file")
	}
	if err = json.Unmarshal(contents, &token); err != nil {
		return token, errors.Wrap(err, "failed to unmarshal token")
	}

	return token, nil
}

func parseInsertChain(path string) (types.Chain, error) {
	var chain types.Chain
	contents, err := os.ReadFile(path)
	if err != nil {
		return chain, errors.Wrap(err, "error reading file")
	}
	if err = json.Unmarshal(contents, &chain); err != nil {
		return chain, errors.Wrap(err, "failed to unmarshal chain")
	}

	return chain, nil
}
