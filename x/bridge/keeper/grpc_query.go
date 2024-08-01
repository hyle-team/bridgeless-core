package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GetTokenById(goctx context.Context, req *types.QueryGetTokenById) (*types.QueryGetTokenByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	id, err := strconv.ParseUint(req.Id, 10, 64)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id %s", req.Id)
	}

	token, found := k.GetToken(ctx, id)
	if !found {
		return nil, types.ErrTokenNotFound
	}

	return &types.QueryGetTokenByIdResponse{
		Token: &token,
	}, nil
}

func (k Keeper) GetTokens(goctx context.Context, req *types.QueryGetTokens) (*types.QueryGetTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	tokens, page, err := k.GetTokensWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryGetTokensResponse{
		Tokens:     tokens,
		Pagination: page,
	}, nil
}

func (k Keeper) GetTokenPairForDstChain(goctx context.Context, req *types.QueryGetTokenPair) (*types.QueryGetTokenPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)

	token, found := k.GetTokenPair(ctx, req.SrcChain, req.SrcAddress, req.DstChain)
	if !found {
		return nil, types.ErrTokenPairsNotFound
	}

	return &types.QueryGetTokenPairResponse{
		Info: token,
	}, nil
}

func (k Keeper) GetAllTokenPairs(goctx context.Context, req *types.QueryGetTokenPairs) (*types.QueryGetTokenPairsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	pairs := k.GetTokenPairs(ctx, req.SrcChain, req.SrcAddress)

	return &types.QueryGetTokenPairsResponse{
		Pairs: pairs,
	}, nil
}

func (k Keeper) GetChainById(goctx context.Context, req *types.QueryGetChainById) (*types.QueryGetChainByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)

	token, found := k.GetChain(ctx, req.Id)
	if !found {
		return nil, types.ErrTokenNotFound
	}

	return &types.QueryGetChainByIdResponse{
		Chain: &token,
	}, nil
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
	return &types.QueryGetChainsResponse{
		Chains:     chains,
		Pagination: page,
	}, nil
}
