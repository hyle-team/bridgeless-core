package bridge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/x/bridge/keeper"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	// this line is used by starport scaffolding # handler/msgServer

	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgSetChain:
			res, err := msgServer.SetAndUpdateChain(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetToken:
			res, err := msgServer.SetAndUpdateToken(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveChain:
			res, err := msgServer.RemoveChainInfo(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveToken:
			res, err := msgServer.RemoveTokenInfo(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetPair:
			res, err := msgServer.SetPairInfo(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemovePair:
			res, err := msgServer.RemovePairInfo(ctx, msg)
			return sdk.WrapServiceResult(ctx, res, err)
		}

		return nil, nil
	}

	// this line is used by starport scaffolding # 1
}
