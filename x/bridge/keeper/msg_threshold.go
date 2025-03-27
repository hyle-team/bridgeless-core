package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (m msgServer) SetStakeThreshold(goCtx context.Context, msg *types.MsgSetStakeThreshold) (*types.MsgSetThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if ok := m.Keeper.IsAdmin(msg.Creator, ctx); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(ctx)
	params.StakeThreshold = msg.Amount
	m.Keeper.SetParams(ctx, params)

	return &types.MsgSetThresholdResponse{}, nil
}

func (m msgServer) SetTssThreshold(goCtx context.Context, msg *types.MsgSetTssThreshold) (*types.MsgSetThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if ok := m.Keeper.IsAdmin(msg.Creator, ctx); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(ctx)
	params.TssThreshold = msg.Amount
	m.Keeper.SetParams(ctx, params)

	return &types.MsgSetThresholdResponse{}, nil
}
