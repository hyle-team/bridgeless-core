package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v12/x/accumulator/types"
	"time"
)

type AdminVesting interface {
	Unlock() error
	UnlockImmediately(amount int64) error
}

type BaseAdminVesting struct {
	address         string
	vesting         *types.VestingInfo
	keeper          Keeper
	lastVestingTime time.Time
}

func NewBaseAdminVesting(keeper Keeper, address string, lastVestingTime time.Time, vesting *types.VestingInfo) BaseAdminVesting {
	return BaseAdminVesting{
		address:         address,
		vesting:         vesting,
		keeper:          keeper,
		lastVestingTime: lastVestingTime,
	}
}

func (av BaseAdminVesting) UnlockImmediately(ctx sdk.Context, amount int64) error {
	fmt.Println("UnlockImmediately")

	fmt.Println("ctx: ", ctx)
	fmt.Println("av: ", av)
	fmt.Println("av.address: ", av.address)
	fmt.Println("amount: ", amount)

	err := av.keeper.sendToAddress(ctx, types.ModuleName, av.address, amount)
	if err != nil {
		return errors.Wrap(err, "failed to unlock money")
	}

	return nil
}

func (av BaseAdminVesting) Unlock() (*time.Time, error) {
	currentTime := time.Now()
	if currentTime.Add(time.Duration(av.vesting.StepTime)*time.Minute).Unix() < av.lastVestingTime.Unix() {
		return nil, nil
	}

	err := av.keeper.sendToAddress(sdk.UnwrapSDKContext(context.Background()), types.ModuleName, av.address, av.vesting.VestingStepAmount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unlock money")
	}

	return &currentTime, nil
}
