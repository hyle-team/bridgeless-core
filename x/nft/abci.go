package nft

import (
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hyle-team/bridgeless-core/x/nft/keeper"
	"time"
)

// update vesting state for each nft
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	for _, nft := range k.GetAllNFT(ctx) {
		if nft.LastVestingTime-ctx.BlockTime().Unix() < nft.VestingPeriod {
			return
		}

		if nft.VestingCounter >= nft.VestingPeriodsCount {
			return
		}

		nft.AvalableToWithdraw.Add(sdk.NewCoin(sdk.NativeToken, sdk.NewInt(nft.VestingPeriod).Mul(nft.RewardPerPeriod.Amount)))
		nft.VestingCounter++
		nft.LastVestingTime = ctx.BlockTime().Unix()

		k.SetNFT(ctx, nft)
	}

}
