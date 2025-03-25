package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (m msgServer) SetParties(goCtx context.Context, msg *types.MsgSetParties) (*types.MsgSetPartiesResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if ok := m.Keeper.IsAdmin(msg.Creator, sdk.UnwrapSDKContext(goCtx)); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(sdk.UnwrapSDKContext(goCtx))
	params.Parties = msg.Parties
	m.SetParams(sdk.UnwrapSDKContext(goCtx), params)

	return &types.MsgSetPartiesResponse{}, nil
}

func (m msgServer) SetNewbieParties(goCtx context.Context, msg *types.MsgSetParties) (*types.MsgSetPartiesResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if ok := m.Keeper.IsAdmin(msg.Creator, sdk.UnwrapSDKContext(goCtx)); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(sdk.UnwrapSDKContext(goCtx))
	params.Newbies = msg.Parties
	m.SetParams(sdk.UnwrapSDKContext(goCtx), params)

	return &types.MsgSetPartiesResponse{}, nil
}

func (m msgServer) SetGoodbyeParties(goCtx context.Context, msg *types.MsgSetParties) (*types.MsgSetPartiesResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if ok := m.Keeper.IsAdmin(msg.Creator, sdk.UnwrapSDKContext(goCtx)); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(sdk.UnwrapSDKContext(goCtx))
	params.GoodbyeList = msg.Parties
	m.SetParams(sdk.UnwrapSDKContext(goCtx), params)

	return &types.MsgSetPartiesResponse{}, nil
}

func (m msgServer) SetBlacklistParties(goCtx context.Context, msg *types.MsgSetParties) (*types.MsgSetPartiesResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if ok := m.Keeper.IsAdmin(msg.Creator, sdk.UnwrapSDKContext(goCtx)); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(sdk.UnwrapSDKContext(goCtx))
	params.Blacklist = msg.Parties
	m.SetParams(sdk.UnwrapSDKContext(goCtx), params)

	return &types.MsgSetPartiesResponse{}, nil
}
