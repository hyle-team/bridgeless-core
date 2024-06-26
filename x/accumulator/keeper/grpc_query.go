package keeper

import (
	"github.com/evmos/evmos/v18/x/accumulator/types"
)

var _ types.QueryServer = BaseKeeper{}
