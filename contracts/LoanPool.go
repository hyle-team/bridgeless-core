package contracts

import (
	// embed compiled smart contract
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// HexString is a byte array that serializes to hex
type HexString []byte

// MarshalJSON serializes ByteArray to hex
func (s HexString) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%x", string(s)))
}

// UnmarshalJSON deserializes ByteArray to hex
func (s *HexString) UnmarshalJSON(data []byte) error {
	var x string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	str, err := hex.DecodeString(x)
	if err != nil {
		return err
	}
	*s = str
	return nil
}

// CompiledContract contains compiled bytecode and abi
type CompiledContract struct {
	ABI abi.ABI
	Bin HexString
}

type jsonCompiledContract struct {
	ABI string
	Bin HexString
}

// MarshalJSON serializes ByteArray to hex
func (s CompiledContract) MarshalJSON() ([]byte, error) {
	abi1, err := json.Marshal(s.ABI)
	if err != nil {
		return nil, err
	}
	return json.Marshal(jsonCompiledContract{ABI: string(abi1), Bin: s.Bin})
}

// UnmarshalJSON deserializes ByteArray to hex
func (s *CompiledContract) UnmarshalJSON(data []byte) error {
	var x jsonCompiledContract
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	s.Bin = x.Bin
	if err := json.Unmarshal([]byte(x.ABI), &s.ABI); err != nil {
		return fmt.Errorf("failed to unmarshal ABI: %w", err)
	}

	return nil
}

var (
	//go:embed compiled_contracts/loancontract.json
	loancontractJSON []byte

	// LoanContract is the compiled test Loan contract
	LoanContract        CompiledContract
	LoanContractAddress common.Address
)

func init() {
	LoanContractAddress = common.HexToAddress("0x7DFE532267a049fB3C5e27E8E2E13B06397f8dD8") // TODO set address

	err := json.Unmarshal(loancontractJSON, &LoanContract)
	if err != nil {
		panic(err)
	}

	if len(LoanContract.Bin) == 0 {
		panic("load contract failed")
	}
}
