package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v18/utils"
)

func (k BaseKeeper) DistributeTokens(ctx sdk.Context, amount int64, moduleNameFrom, moduleNameTo string, address *string) error {
	if err := k.validateBalance(ctx, moduleNameFrom, amount); err != nil {
		return errors.Wrap(err, "failed to validate balance")
	}

	switch moduleNameFrom {
	case "distribution":
		if address == nil || *address == "" {
			return errors.Wrap(fmt.Errorf("module address cannot be empty"), "from distribution: DistributeTokens")
		}
		err := k.sendToAddress(ctx, moduleNameFrom, *address, amount)
		if err != nil {
			return errors.Wrap(err, "failed to send tokens form distribution")
		}
		break
	default:
		break
	}

	return nil
}

func (k BaseKeeper) sendToModuleAddress(ctx sdk.Context, moduleNameFrom, moduleNameTo string, amount int64) error {
	moduleInfoFrom := k.GetModuleInfo(ctx, moduleNameFrom)
	moduleInfoTo := k.GetModuleInfo(ctx, moduleNameTo)

	err := k.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		moduleInfoFrom.Address,
		moduleInfoTo.Address,
		utils.NewnNativeTokens(amount),
	)

	if err != nil {
		return errors.Wrap(err, "sending native coins to address")
	}

	return nil
}

// TODO use enum  instead of module name or
func (k BaseKeeper) sendToAddress(ctx sdk.Context, moduleNameFrom, address string, amount int64) error {
	moduleInfoFrom := k.GetModuleInfo(ctx, moduleNameFrom)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		moduleInfoFrom.Address,
		sdk.AccAddress(address),
		utils.NewnNativeTokens(amount),
	)

	if err != nil {
		return errors.Wrap(err, "sending native coins to address")
	}

	return nil
}
