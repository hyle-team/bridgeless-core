package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) RemoveTokenInfo(goctx context.Context, msg *types.MsgRemoveToken) (*types.MsgRemoveTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if k.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}
	return nil, nil

}

func (k msgServer) SetAndUpdateToken(goctx context.Context, msg *types.MsgSetToken) (*types.MsgSetTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if k.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	k.SetToken(ctx, *msg.Token)
	return &types.MsgSetTokenResponse{}, nil

}

func (k msgServer) RemoveChainInfo(goctx context.Context, msg *types.MsgRemoveChain) (*types.MsgRemoveChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if k.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	k.RemoveChain(ctx, msg.ChainId)

	return &types.MsgRemoveChainResponse{}, nil

}

func (k msgServer) SetAndUpdateChain(goctx context.Context, msg *types.MsgSetChain) (*types.MsgSetChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if k.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	k.SetChain(ctx, *msg.Chain)

	return &types.MsgSetChainResponse{}, nil
}
