package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetChainById(goctx context.Context, req *types.QueryGetChainById) (*types.QueryGetChainByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)

	token, found := k.GetChain(ctx, req.Id)
	if !found {
		return nil, types.ErrChainNotFound
	}

	return &types.QueryGetChainByIdResponse{Chain: token}, nil
}

func (k Keeper) GetChains(goctx context.Context, req *types.QueryGetChains) (*types.QueryGetChainsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	chains, page, err := k.GetChainsWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, types.ErrChainNotFound
	}

	return &types.QueryGetChainsResponse{Chains: chains, Pagination: page}, nil
}
