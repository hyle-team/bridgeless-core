package nft

import (
	"github.com/hyle-team/bridgeless-core/x/nft/keeper"
	"github.com/hyle-team/bridgeless-core/x/nft/types"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// update vesting state for each nft
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	for _, nft := range k.GetNFTs() {
		nft.UpdateVesting(ctx)
	}
}
