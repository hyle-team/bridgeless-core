package keeper

import (
	"github.com/hyle-team/bridgeless-core/x/nft/types"
)

var _ types.QueryServer = Keeper{}
