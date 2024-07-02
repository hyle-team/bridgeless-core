package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v12/utils"
)

func (k BaseKeeper) MintTokens(ctx sdk.Context, amount int64, moduleName string) error {
	fmt.Println("mint tokens")
	if err := k.bankKeeper.MintCoins(ctx, moduleName, utils.NewNativeTokens(amount)); err != nil {
		return errors.Wrap(err, "mint native tokens")
	}

	return nil
}

func (k BaseKeeper) sendToModuleAddress(ctx sdk.Context, moduleNameFrom, moduleNameTo string, amount int64) error {
	moduleInfoFrom := k.GetModuleInfo(ctx, moduleNameFrom)
	if moduleInfoFrom == nil {
		fmt.Println("no moduleFrom  info found")
		return nil
	}
	moduleInfoTo := k.GetModuleInfo(ctx, moduleNameTo)
	if moduleInfoTo == nil {
		fmt.Println("no moduleTo info found")
		return nil
	}

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
	fmt.Println("sendToAddress")

	moduleInfoFrom := k.GetModuleInfo(ctx, moduleNameFrom)
	if moduleInfoFrom == nil {
		return fmt.Errorf("no module info found")
	}

	fmt.Println(ctx,
		moduleInfoFrom.Address,
		sdk.AccAddress(address),
		utils.NewNativeTokens(amount))

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
