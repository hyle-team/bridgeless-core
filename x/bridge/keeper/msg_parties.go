package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
)

func (m msgServer) SetParties(goCtx context.Context, msg *types.MsgSetPartiesList) (*types.MsgSetPartiesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if ok := m.Keeper.IsAdmin(msg.Creator, ctx); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(ctx)
	params.Parties = msg.Parties
	m.SetParams(ctx, params)

	return &types.MsgSetPartiesResponse{}, nil
}

func (m msgServer) SetNewbieParties(goCtx context.Context, msg *types.MsgSetNewbiesList) (*types.MsgSetPartiesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if ok := m.Keeper.IsAdmin(msg.Creator, ctx); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(ctx)
	params.Newbies = msg.Parties
	m.SetParams(ctx, params)

	return &types.MsgSetPartiesResponse{}, nil
}

func (m msgServer) SetGoodbyeParties(goCtx context.Context, msg *types.MsgSetGoodbyePartiesList) (*types.MsgSetPartiesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if ok := m.Keeper.IsAdmin(msg.Creator, ctx); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(ctx)
	params.GoodbyeList = msg.Parties
	m.SetParams(ctx, params)

	return &types.MsgSetPartiesResponse{}, nil
}

func (m msgServer) SetBlacklistParties(goCtx context.Context, msg *types.MsgSetBlacklistPartiesList) (*types.MsgSetPartiesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	if ok := m.Keeper.IsAdmin(msg.Creator, ctx); !ok {
		return nil, types.ErrPermissionDenied
	}

	params := m.Keeper.GetParams(ctx)
	params.Blacklist = msg.Parties
	m.SetParams(ctx, params)

	return &types.MsgSetPartiesResponse{}, nil
}
