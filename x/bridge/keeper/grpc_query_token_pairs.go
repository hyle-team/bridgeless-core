package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k queryServer) GetTokenPair(goctx context.Context, req *types.QueryGetTokenPair) (*types.QueryGetTokenPairResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goctx)
	pair, found := k.GetDstToken(ctx, req.SrcAddress, req.SrcChain, req.DstChain)
	if !found {
		return nil, types.ErrTokenPairNotFound
	}

	return &types.QueryGetTokenPairResponse{Info: pair}, nil
}
