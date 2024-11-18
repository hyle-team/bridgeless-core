package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetAllPositions(goCtx context.Context, req *types.QueryPositions) (*types.QueryPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	positions, pagination, err := k.GetPositionsWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, err
	}

	return &types.QueryPositionsResponse{Positions: positions, Pagination: pagination}, nil
}

func (k Keeper) GetPositionByAddress(goCtx context.Context, req *types.QueryPositionByAddress) (*types.QueryPositionByAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	position, ok := k.GetPosition(ctx, req.Address)
	if !ok {
		return nil, errors.New("position not found")
	}

	return &types.QueryPositionByAddressResponse{Position: position}, nil
}
