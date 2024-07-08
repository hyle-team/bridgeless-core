package keeper

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/hyle-team/bridgeless-core/x/nft/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper    types.BankKeeper
		stakingKeeper types.StakingKeeper
		nfts          map[string]NFT
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,

) *Keeper {
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		nfts:          make(map[string]NFT),
		bankKeeper:    bankKeeper,
		stakingKeeper: stakingKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) DistributeToAddress(ctx sdk.Context, amount sdk.Coins, owner sdk.AccAddress, nftAddress sdk.AccAddress) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, nftAddress, types.ModuleName, amount)
	if err != nil {
		err = errors.Wrap(err, "failed to distribute tokens to nft module")
		k.Logger(ctx).Error(err.Error())
		return err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, owner, amount)
	if err != nil {
		err = errors.Wrap(err, "failed to distribute tokens to nft module")
		k.Logger(ctx).Error(err.Error())
		return err
	}

	return nil
}

func (k Keeper) Mint(rawNft types.NFT) NFT {
	return NewNft(
		k,
		rawNft.Owner,
		rawNft.Uri,
		rawNft.RewardPerPeriod,
		rawNft.VestingPeriodsCount,
		rawNft.VestingPeriod,
		rawNft.Address,
	)
}

func (k Keeper) AppendNFT(nft NFT) {
	k.nfts[nft.GetAddress()] = nft
}

func (k Keeper) GetNFTs() map[string]NFT {
	return k.nfts
}

func (k Keeper) Undelegate(goctx context.Context, request *types.QueryUndelegateRequest) (*types.QueryUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	validator, err := sdk.ValAddressFromBech32(request.Address)
	if err != nil {
		return new(types.QueryUndelegateResponse), err
	}
	return new(types.QueryUndelegateResponse), k.nfts[request.Address].Unstake(ctx, validator)
}

func (k Keeper) Delegate(goctx context.Context, request *types.QueryDelegateRequest) (*types.QueryDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	validator, err := sdk.ValAddressFromBech32(request.Address)
	if err != nil {
		return new(types.QueryDelegateResponse), err
	}
	return new(types.QueryDelegateResponse), k.nfts[request.Address].Stake(ctx, validator)

}

func (k Keeper) Send(goctx context.Context, request *types.QuerySendRequest) (*types.QuerySendResponse, error) {
	k.nfts[request.Address].Send(request.Recipient)
	return new(types.QuerySendResponse), nil

}

func (k Keeper) Withdraw(goctx context.Context, request *types.QueryWithdrawalRequest) (*types.QueryWithdrawalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	return new(types.QueryWithdrawalResponse), k.nfts[request.Address].Withdraw(ctx)

}
