// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package testdata

import (
	contractutils "github.com/hyle-team/bridgeless-core/contracts/utils"
	evmtypes "github.com/hyle-team/bridgeless-core/x/evm/types"
)

func LoadERC20AllowanceCaller() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20AllowanceCaller.json")
}
