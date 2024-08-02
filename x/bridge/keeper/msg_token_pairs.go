package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func (m msgServer) AddTokenInfo(goCtx context.Context, msg *types.MsgAddTokenInfo) (*types.MsgAddTokenInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	token, found := m.GetToken(ctx, msg.Info.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "token not found")
	}

	for _, info := range token.Info {
		if info.ChainId == msg.Info.ChainId {
			return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "token info already exists")
		}
	}

	token.Info = append(token.Info, msg.Info)
	m.SetToken(ctx, token)

	// mapping src chain -> dest chain
	m.setTokenPairs(ctx, msg.Info, token.Info)
	for _, info := range token.Info {
		// reverse mapping dest chain -> src chain
		m.setTokenPairs(ctx, info, []types.TokenInfo{msg.Info})
	}

	return &types.MsgAddTokenInfoResponse{}, nil
}

func (m msgServer) RemoveTokenInfo(goCtx context.Context, msg *types.MsgRemoveTokenInfo) (*types.MsgRemoveTokenInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	token, found := m.GetToken(ctx, msg.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "token not found")
	}

	var idx = -1
	for i, info := range token.Info {
		if info.ChainId == msg.ChainId {
			idx = i
			break
		}
	}
	if idx == -1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "token info not found")
	}

	// mapping src chain -> dest chain
	m.removeTokenPairs(ctx, token.Info[idx], token.Info)
	for _, info := range token.Info {
		// reverse mapping dest chain -> src chain
		m.removeTokenPairs(ctx, info, []types.TokenInfo{token.Info[idx]})
	}

	token.Info = append(token.Info[:idx], token.Info[idx+1:]...)
	m.SetToken(ctx, token)

	return &types.MsgRemoveTokenInfoResponse{}, nil
}
