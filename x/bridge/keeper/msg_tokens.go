package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func (m msgServer) InsertToken(goCtx context.Context, msg *types.MsgInsertToken) (*types.MsgInsertTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	_, found := m.GetToken(ctx, msg.Token.Id)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrConflict, "token already exists")
	}

	m.SetToken(ctx, msg.Token)
	for _, pair := range msg.Token.Info {
		m.SetTokenPairs(ctx, pair, msg.Token.Info...)
	}

	return &types.MsgInsertTokenResponse{}, nil
}

func (m msgServer) UpdateToken(goCtx context.Context, msg *types.MsgUpdateToken) (*types.MsgUpdateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	token, found := m.GetToken(ctx, msg.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "token not found")
	}

	token.Metadata = msg.Metadata
	m.SetToken(ctx, token)

	return &types.MsgUpdateTokenResponse{}, nil
}

func (m msgServer) DeleteToken(goCtx context.Context, msg *types.MsgDeleteToken) (*types.MsgDeleteTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Creator != m.GetParams(ctx).ModuleAdmin {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "msg sender is not module admin")
	}

	token, found := m.GetToken(ctx, msg.TokenId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotFound, "token not found")
	}

	m.RemoveToken(ctx, msg.TokenId)
	for _, pair := range token.Info {
		m.RemoveTokenPairs(ctx, pair, token.Info...)
	}

	return &types.MsgDeleteTokenResponse{}, nil
}
