package keeper

import (
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/types"
)

var _ types.QueryServer = Keeper{}
