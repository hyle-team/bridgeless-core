package keeper

import (
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
