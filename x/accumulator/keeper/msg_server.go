package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v18/x/accumulator/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) Reset() {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) String() string {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) ProtoMessage() {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m msgServer) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServiceServer {
	return &msgServer{keeper}
}

var _ types.MsgServiceServer = msgServer{}
var _ sdk.Msg = msgServer{}

func (m msgServer) DistributeTokens(ctx context.Context, request *types.DistributeTokensRequest) (*types.DistributeTokensResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if err := m.validateBalance(sdkCtx, request.ModuleNameFrom, request.Amount); err != nil {
		return nil, errors.Wrap(err, "failed to validate balance")
	}

	// TODO use enum for module names instead of string
	switch request.ModuleNameFrom {
	case "distribution":
		if request.Address == "" {
			return nil, errors.Wrap(fmt.Errorf("module address cannot be empty"), "from distribution: DistributeTokens")
		}
		err := m.sendToAddress(sdkCtx, request.ModuleNameFrom, request.Address, request.Amount)
		if err != nil {
			return nil, errors.Wrap(err, "failed to send tokens form distribution")
		}
		break
	default:
		break
	}

	return nil, nil
}
