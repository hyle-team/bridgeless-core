package keeper

import (
	"context"
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v18/x/accumulator/types"
	"time"
)

const accumulator = "accumulator"

type AdminVesting interface {
	Unlock() error
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

func (av BaseAdminVesting) UnlockImmediately(amount int64) error {
	err := av.keeper.sendToAddress(sdk.UnwrapSDKContext(context.Background()), accumulator, av.address, amount)
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

	err := av.keeper.sendToAddress(sdk.UnwrapSDKContext(context.Background()), accumulator, av.address, av.vesting.VestingStepAmount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unlock money")
	}

	return &currentTime, nil
}
