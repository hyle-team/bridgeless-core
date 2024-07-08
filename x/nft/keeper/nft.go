package keeper

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/hyle-team/bridgeless-core/x/nft/types"
	"time"
)

type NFT interface {
	URI() string
	Withdraw(ctx sdk.Context) error
	CheckIsDelegated(ctx sdk.Context) (bool, error)
	Stake(ctx sdk.Context, validatorAddress sdk.ValAddress) error
	Unstake(ctx sdk.Context, validatorAddress sdk.ValAddress) error
	GetStakedBalance(ctx sdk.Context, delegatorAddr sdk.AccAddress) (sdk.Coin, error)
	UpdateVesting(ctx sdk.Context)
	GetAddress() string
	Send(recipient string)
}

type BaseNFT struct {
	keeper           Keeper
	lastVestingTime  time.Time
	availableBalance sdk.Coin
	vestingCounter   int64
	types.NFT
}

func NewNft(keeper Keeper, owner string, uri string, rewardPerPeriod sdk.Coin, vestingCount, vestingPeriod int64, address string) NFT {
	nft := types.NFT{
		Owner:               owner,
		Uri:                 uri,
		RewardPerPeriod:     rewardPerPeriod,
		VestingPeriodsCount: vestingCount,
		VestingPeriod:       vestingPeriod,
		Address:             address,
	}

	return &BaseNFT{
		keeper:           keeper,
		lastVestingTime:  time.Now(),
		availableBalance: sdk.NewCoin(sdk.NativeToken, sdk.NewInt(0)),
		NFT:              nft,
	}
}

func (n BaseNFT) URI() string {
	return n.Uri
}

func (n BaseNFT) CheckIsDelegated(ctx sdk.Context) (bool, error) {
	address, err := sdk.AccAddressFromBech32(n.Address)
	if err != nil {
		err = errors.Wrap(err, "failed to decode address")
	}

	delegators := n.keeper.stakingKeeper.GetAllDelegatorDelegations(ctx, address)

	return len(delegators) > 0, nil
}

func (n BaseNFT) Stake(ctx sdk.Context, validatorAddress sdk.ValAddress) error {
	address, err := sdk.AccAddressFromBech32(n.Address)
	if err != nil {
		n.keeper.Logger(ctx).Error(err.Error())
		return errors.Wrap(err, "failed to decode address")
	}
	balances := n.keeper.bankKeeper.GetAllBalances(ctx, address)

	if balances.AmountOf(sdk.NativeToken).IsZero() {
		return types.ErrInvalidAmount
	}

	validator, found := n.keeper.stakingKeeper.GetValidator(ctx, validatorAddress)
	if !found {
		return types.ErrValidatorNotFound
	}

	_, err = n.keeper.stakingKeeper.Delegate(ctx, address, balances.AmountOf(sdk.NativeToken), stakingtypes.Unbonded, validator, true)
	if err != nil {
		return types.ErrFailedToDelegate
	}

	err = n.keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, address, stakingtypes.BondedPoolName, balances)
	if err != nil {
		return types.ErrFailedToSendTokenAmount
	}

	return nil
}

func (n BaseNFT) Unstake(ctx sdk.Context, validatorAddress sdk.ValAddress) error {
	address, err := sdk.AccAddressFromBech32(n.Address)
	if err != nil {
		return errors.Wrap(err, "failed to decode address")
	}

	_, err = n.keeper.stakingKeeper.Undelegate(ctx, address, validatorAddress, sdk.ZeroDec())
	if err != nil {
		err = errors.Wrap(err, "failed to undelegate tokens")
		n.keeper.Logger(ctx).Error(err.Error())
		return err
	}
	return nil
}

func (n BaseNFT) Withdraw(ctx sdk.Context) error {
	isDelegated, err := n.CheckIsDelegated(ctx)
	if !isDelegated {
		return types.ErrTokenIsDelegated
	}

	if n.availableBalance.IsZero() {
		return types.ErrInvalidAmount
	}

	nftAddress, err := sdk.AccAddressFromBech32(n.Address)
	if err != nil {
		err = errors.Wrap(err, "failed to decode nft address")
		n.keeper.Logger(ctx).Error(err.Error())
		return err
	}

	ownerAddress, err := sdk.AccAddressFromBech32(n.Owner)
	if err != nil {
		err = errors.Wrap(err, "failed to decode owner address")
		n.keeper.Logger(ctx).Error(err.Error())
		return err
	}

	if err = n.keeper.DistributeToAddress(ctx, sdk.NewCoins(n.availableBalance), ownerAddress, nftAddress); err != nil {
		return errors.Wrap(err, "failed to distribute nft to address")
	}

	return nil
}

func (n BaseNFT) GetStakedBalance(ctx sdk.Context, delegatorAddr sdk.AccAddress) (sdk.Coin, error) {
	delegations := n.keeper.stakingKeeper.GetAllDelegatorDelegations(ctx, delegatorAddr)

	totalStaked := sdk.NewCoin(sdk.DefaultBondDenom, sdk.ZeroInt())
	for _, delegation := range delegations {
		validator, found := n.keeper.stakingKeeper.GetValidator(ctx, delegation.GetValidatorAddr())
		if !found {
			return sdk.Coin{}, types.ErrValidatorNotFound
		}
		shares := delegation.Shares
		tokens := validator.TokensFromSharesTruncated(shares)
		totalStaked = totalStaked.Add(sdk.NewCoin(sdk.DefaultBondDenom, tokens.TruncateInt()))
	}

	return totalStaked, nil
}

func (n BaseNFT) UpdateVesting(ctx sdk.Context) {
	if n.lastVestingTime.Unix()-ctx.BlockTime().Unix() < n.VestingPeriod {
		return
	}

	if n.vestingCounter >= n.VestingPeriodsCount {
		return
	}

	n.availableBalance.Add(sdk.NewCoin(sdk.NativeToken, sdk.NewInt(n.VestingPeriod).Mul(n.RewardPerPeriod.Amount)))
	n.vestingCounter++
	n.lastVestingTime = ctx.BlockTime()
}

func (n BaseNFT) GetAddress() string {
	return n.Address
}

func (n BaseNFT) Send(recipient string) {
	n.Owner = recipient
}
