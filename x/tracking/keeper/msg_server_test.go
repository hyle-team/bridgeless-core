package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/hyle-team/bridgeless-core/x/tracking/types"
    "github.com/hyle-team/bridgeless-core/x/tracking/keeper"
    keepertest "github.com/hyle-team/bridgeless-core/testutil/keeper"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TrackingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
