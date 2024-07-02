package keeper

import (
	"github.com/evmos/evmos/v12/x/nft/types"
)

var _ types.QueryServer = Keeper{}
