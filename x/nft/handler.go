package nft

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/nft/keeper"
	"github.com/hyle-team/bridgeless-core/x/nft/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	// this line is used by starport scaffolding # handler/msgServer

	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgUndelegate:
			res, err := msgServer.Undelegate(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDelegate:
			res, err := msgServer.Delegate(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSend:
			res, err := msgServer.Send(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgWithdrawal:
			res, err := msgServer.Withdraw(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRedelegate:
			res, err := msgServer.Redelegate(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		}

		return nil, nil
	}

	// this line is used by starport scaffolding # 1
}
