package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) GetTokenById(goctx context.Context, req *types.QueryGetTokenById) (*types.QueryGetTokenByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	token, found := k.GetToken(ctx, req.Id)
	if !found {
		return nil, types.ErrTokenNotFound
	}

	return &types.QueryGetTokenByIdResponse{Token: token}, nil
}

func (k queryServer) GetTokens(goctx context.Context, req *types.QueryGetTokens) (*types.QueryGetTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	tokens, page, err := k.GetTokensWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetTokensResponse{Tokens: tokens, Pagination: page}, nil
}
