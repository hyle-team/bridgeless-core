package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (m msgServer) SetTssThreshold(goCtx context.Context, msg *types.MsgSetTssThreshold) (*types.MsgSetTssThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := m.Keeper.GetParams(ctx)
	params.TssThreshold = msg.Threshold
	m.Keeper.SetParams(ctx, params)

	return &types.MsgSetTssThresholdResponse{}, nil
}
