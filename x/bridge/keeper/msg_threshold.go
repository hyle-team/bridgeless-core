package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (m msgServer) SetStakeThreshold(goCtx context.Context, msg *types.MsgSetStakeThreshold) (*types.MsgSetThresholdResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if ok := m.Keeper.IsAdmin(msg.Creator, sdk.UnwrapSDKContext(goCtx)); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(sdk.UnwrapSDKContext(goCtx))
	params.StakeThreshold = msg.Amount
	m.Keeper.SetParams(sdk.UnwrapSDKContext(goCtx), params)

	return &types.MsgSetThresholdResponse{}, nil
}

func (m msgServer) SetTssThreshold(goCtx context.Context, msg *types.MsgSetTssThreshold) (*types.MsgSetThresholdResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if ok := m.Keeper.IsAdmin(msg.Creator, sdk.UnwrapSDKContext(goCtx)); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(sdk.UnwrapSDKContext(goCtx))
	params.TssThreshold = msg.Amount
	m.Keeper.SetParams(sdk.UnwrapSDKContext(goCtx), params)

	return &types.MsgSetThresholdResponse{}, nil
}
