package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/hyle-team/bridgeless-core/v12/testutil/keeper"
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/keeper"
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TrackingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
