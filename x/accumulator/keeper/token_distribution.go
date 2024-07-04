package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k BaseKeeper) SendToModuleAddress(ctx sdk.Context, moduleNameFrom, moduleNameTo string, amount sdk.Coins) error {
	moduleInfoFrom := k.GetModuleInfo(ctx, moduleNameFrom)
	if moduleInfoFrom == nil {
		k.Logger(ctx).Error("no moduleFrom info found")
		return nil
	}
	moduleInfoTo := k.GetModuleInfo(ctx, moduleNameTo)
	if moduleInfoTo == nil {
		err := fmt.Errorf("no moduleTo info found")
		k.Logger(ctx).Error(err.Error())
		return err
	}

	moduleAddrFrom, err := sdk.AccAddressFromBech32(moduleInfoFrom.Address)
	if err != nil {
		err = errors.Wrap(err, "failed to parse address")
		k.Logger(ctx).Error(err.Error())
		return err
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		moduleAddrFrom,
		moduleNameFrom,
		amount,
	)

	if err != nil {
		return errors.Wrap(err, "sending native coins to module")
	}

	return nil
}

func (k BaseKeeper) SendFromModuleToAddress(ctx sdk.Context, moduleNameFrom string, address sdk.AccAddress, amount sdk.Coins) error {
	moduleInfo := k.GetModuleInfo(ctx, moduleNameFrom)
	if moduleInfo == nil {
		err := fmt.Errorf("no module info found")
		k.Logger(ctx).Error(err.Error())
		return err
	}
	fmt.Println("address", address)

	err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		moduleNameFrom,
		address,
		amount,
	)

	if err != nil {
		return errors.Wrap(err, "sending native coins from module to address")
	}

	return nil
}

func (k BaseKeeper) SendFromAddressToAddress(ctx sdk.Context, moduleNameFrom, address string, amount sdk.Coins) error {
	fmt.Println("sendToAddress")

	moduleInfo := k.GetModuleInfo(ctx, moduleNameFrom)
	if moduleInfo == nil {
		err := fmt.Errorf("no module info found")
		k.Logger(ctx).Error(err.Error())
		return err
	}

	fmt.Println("moduleInfo: ", moduleInfo)
	recipientAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		err = errors.Wrap(err, "failed to parse address")
		k.Logger(ctx).Error(err.Error())
		return err
	}

	fmt.Println("recipientAddr: ", recipientAddr.String())

	moduleAddress, err := sdk.AccAddressFromBech32(moduleInfo.Address)
	if err != nil {
		err = errors.Wrap(err, "failed to parse address")
		k.Logger(ctx).Error(err.Error())
		return err
	}

	fmt.Println("moduleAddress: ", moduleAddress)

	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		moduleAddress,
		moduleNameFrom,
		amount,
	)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return err
	}

	fmt.Println("SendCoinsFromAccountToModule done")

	for _, coin := range amount {
		if !k.bankKeeper.HasBalance(ctx, moduleAddress, coin) {
			err = fmt.Errorf("module %s does not have enough balance to send %s", moduleNameFrom, amount)
			k.Logger(ctx).Error(err.Error())
			return err
		}
	}

	fmt.Println("Sending coins...")

	fmt.Println("module balance ", k.bankKeeper.GetBalance(ctx, authtypes.NewModuleAddress(moduleNameFrom), sdk.NativeToken))
	fmt.Println("amount: ", amount)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		moduleNameFrom,
		recipientAddr,
		amount,
	)

	if err != nil {
		return errors.Wrap(err, "sending native coins to address")
	}

	return nil
}
