package keeper

import (
	"github.com/evmos/evmos/v12/x/accumulator/types"
)

var _ types.QueryServer = BaseKeeper{}
