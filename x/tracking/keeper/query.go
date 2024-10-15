package keeper

import (
	"github.com/hyle-team/bridgeless-core/x/tracking/types"
)

var _ types.QueryServer = Keeper{}
