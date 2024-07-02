package keeper

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v12/utils"
)

func (k BaseKeeper) MintTokens(ctx sdk.Context, amount int64, moduleName string) error {
	if err := k.bankKeeper.MintCoins(ctx, moduleName, utils.NewNativeTokens(amount)); err != nil {
		return errors.Wrap(err, "mint native tokens")
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
		utils.NewNativeTokens(amount),
	)

	if err != nil {
		return errors.Wrap(err, "sending native coins to address")
	}

	return nil
}

func (k BaseKeeper) sendToAddress(ctx sdk.Context, moduleNameFrom, address string, amount int64) error {
	moduleInfoFrom := k.GetModuleInfo(ctx, moduleNameFrom)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		moduleInfoFrom.Address,
		sdk.AccAddress(address),
		utils.NewNativeTokens(amount),
	)

	if err != nil {
		return errors.Wrap(err, "sending native coins to address")
	}

	return nil
}
