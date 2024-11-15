package keeper

import (
	"github.com/hyle-team/bridgeless-core/v12/x/multisig/types"
)

var _ types.QueryServer = Keeper{}
