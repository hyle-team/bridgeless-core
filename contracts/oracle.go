package contracts

import (
	// embed compiled smart contract
	_ "embed"
	"encoding/json"
)

var (
	//go:embed compiled_contracts/Oracle.json
	OracleContractJSON []byte

	// OracleContract is the compiled test Loan contract
	OracleContract CompiledContract
)

func init() {
	err := json.Unmarshal(OracleContractJSON, &OracleContract)
	if err != nil {
		panic(err)
	}

	if len(LoanContract.Bin) == 0 {
		panic("oracle contract failed")
	}
}
