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

func (m msgServer) DeleteToken(goctx context.Context, msg *types.MsgDeleteToken) (*types.MsgDeleteTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if m.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	m.RemoveToken(ctx, msg.TokenId)

	return &types.MsgDeleteTokenResponse{}, nil

}

func (m msgServer) InsertToken(goctx context.Context, msg *types.MsgInsertToken) (*types.MsgInsertTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if m.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	m.SetToken(ctx, *msg.Token)
	return &types.MsgInsertTokenResponse{}, nil

}

func (m msgServer) DeleteChain(goctx context.Context, msg *types.MsgDeleteChain) (*types.MsgDeleteChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if m.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	m.RemoveChain(ctx, msg.ChainId)

	return &types.MsgDeleteChainResponse{}, nil

}

func (m msgServer) InsertChain(goctx context.Context, msg *types.MsgInsertChain) (*types.MsgInsertChainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if m.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	m.SetChain(ctx, *msg.Chain)

	return &types.MsgInsertChainResponse{}, nil
}

func (m msgServer) UpdateToken(goctx context.Context, msg *types.MsgUpdateToken) (*types.MsgUpdateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if m.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	token, found := m.GetToken(ctx, msg.TokenId)
	if !found {
		return nil, types.ErrTokenNotFound
	}

	if msg.Name != "" {
		token.Name = msg.Name
	}

	if msg.Symbol != "" {
		token.Symbol = msg.Symbol
	}

	return &types.MsgUpdateTokenResponse{}, nil
}

func (m msgServer) RemovePairById(goctx context.Context, msg *types.MsgRemovePairById) (*types.MsgRemovePairByIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if m.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	token, found := m.GetToken(ctx, msg.TokenId)
	if !found {
		return nil, types.ErrTokenNotFound
	}
	for _, pair := range token.Info {
		if pair.SourceChain == msg.ChainId {
			err := m.RemoveTokenPair(ctx, pair.SourceChain, pair.DestinationChain, pair.Address)
			if err != nil {
				return nil, err
			}
		}
	}

	return &types.MsgRemovePairByIdResponse{}, nil
}

func (m msgServer) RemovePairByPath(goctx context.Context, msg *types.MsgRemovePairByPath) (*types.MsgRemovePairByPathResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if m.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	err := m.RemoveTokenPair(ctx, msg.SrcChain, msg.DstChain, msg.Address)
	if err != nil {
		return nil, err
	}

	return &types.MsgRemovePairByPathResponse{}, nil
}

func (m msgServer) InsertPairInfo(goctx context.Context, msg *types.MsgSetPair) (*types.MsgSetPairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	if m.GetParams(ctx).EvmAdmin != msg.Creator {
		return nil, types.ErrPermissionDenied
	}

	err := m.SetTokenInfo(ctx, msg.Pair.SourceChain, msg.Pair.DestinationChain, msg.Pair.Address, msg.Pair)
	if err != nil {
		return nil, err
	}

	return &types.MsgSetPairResponse{}, nil
}
