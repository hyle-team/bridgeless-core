// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package testdata

import (
	contractutils "github.com/hyle-team/bridgeless-core/contracts/utils"
	evmtypes "github.com/hyle-team/bridgeless-core/x/evm/types"
)

// LoadMaliciousDelayedContract loads the ERC20MaliciousDelayed contract.
//
// This is an evil token. Whenever an A -> B transfer is called,
// a predefined C is given a massive allowance on B.
func LoadMaliciousDelayedContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20MaliciousDelayed.json")
}
