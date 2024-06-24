package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v18/utils"
	"github.com/evmos/evmos/v18/x/accumulator/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) DistributeTokens(ctx sdk.Context, amount int64, moduleNameFrom, moduleNameTo string, address *string) error {
	switch moduleNameFrom {
	case "distribution":
		if address == nil || *address == "" {
			return errors.Wrap(fmt.Errorf("module address cannot be empty"), "from distribution: DistributeTokens")
		}
		m.sendToAddress(ctx, moduleNameFrom, *address, amount)
		break
	default:
		break
	}

	return nil
}

func (m msgServer) sendToModuleAddress(ctx sdk.Context, moduleNameFrom, moduleNameTo string, amount int64) error {
	moduleInfoFrom := m.GetModuleInfo(moduleNameFrom)
	moduleInfoTo := m.GetModuleInfo(moduleNameTo)

	err := m.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		moduleInfoFrom.address,
		moduleInfoTo.address,
		utils.NewnNativeTokens(amount),
	)

	if err != nil {
		return errors.Wrap(err, "sending native coins to address")
	}

	return nil
}

// TODO use enum  instead of module name or
func (m msgServer) sendToAddress(ctx sdk.Context, moduleNameFrom, address string, amount int64) error {
	moduleInfoFrom := m.GetModuleInfo(moduleNameFrom)
	err := m.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		moduleInfoFrom.address,
		sdk.AccAddress(address),
		utils.NewnNativeTokens(amount),
	)

	if err != nil {
		return errors.Wrap(err, "sending native coins to address")
	}

	return nil
}

var _ types.MsgServer = msgServer{}
