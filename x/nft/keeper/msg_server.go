package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/hyle-team/bridgeless-core/x/nft/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) Undelegate(goctx context.Context, request *types.MsgUndelegate) (*types.MsgUndelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	nft, ok := k.GetNFT(ctx, request.Address)
	if !ok {
		return nil, types.ErrNFTNotFound
	}

	if nft.Owner != request.Creator {
		return nil, types.ErrNFTInvalidOwner
	}

	validator, err := sdk.ValAddressFromBech32(request.Validator)
	if err != nil {
		return nil, err
	}

	if _, found := k.stakingKeeper.GetValidator(ctx, validator); !found {
		return nil, types.ErrValidatorNotFound
	}

	address, err := sdk.AccAddressFromBech32(request.Address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode address")
	}

	_, err = k.stakingKeeper.Undelegate(ctx, address, validator, sdk.NewDecCoinFromCoin(request.Amount).Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to undelegate tokens")
	}

	return new(types.MsgUndelegateResponse), nil
}

func (k msgServer) Delegate(goctx context.Context, request *types.MsgDelegate) (*types.MsgDelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	nft, ok := k.GetNFT(ctx, request.Address)
	if !ok {
		return nil, types.ErrNFTNotFound
	}

	if nft.Owner != request.Creator {
		return nil, types.ErrNFTInvalidOwner
	}

	valAddr, _ := sdk.ValAddressFromBech32(request.Address)
	address, _ := sdk.AccAddressFromBech32(request.Address)

	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return nil, types.ErrValidatorNotFound
	}

	_, err := k.stakingKeeper.Delegate(ctx, address, request.Amount.Amount, stakingtypes.Unbonded, validator, true)
	if err != nil {
		return nil, types.ErrFailedToDelegate
	}

	return new(types.MsgDelegateResponse), nil

}

func (k msgServer) Send(goctx context.Context, request *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	nft, ok := k.GetNFT(ctx, request.Address)
	if !ok {
		return nil, types.ErrNFTNotFound
	}

	if nft.Owner != request.Creator {
		return nil, types.ErrNFTInvalidOwner
	}

	nft.Owner = request.Recipient
	k.SetNFT(ctx, nft)
	return new(types.MsgSendResponse), nil

}

func (k msgServer) Withdraw(goctx context.Context, request *types.MsgWithdrawal) (*types.MsgWithdrawalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	nft, ok := k.GetNFT(ctx, request.Address)
	if !ok {
		return nil, types.ErrNFTNotFound
	}

	if nft.Owner != request.Creator {
		return nil, types.ErrNFTInvalidOwner
	}

	address, _ := sdk.AccAddressFromBech32(request.Address)
	ownerAddress, _ := sdk.AccAddressFromBech32(request.Owner)

	balance := k.bankKeeper.GetAllBalances(ctx, address).AmountOf(sdk.NativeToken)
	amount := sdk.MinInt(nft.AvalableToWithdraw.Amount, balance)

	err := k.DistributeToAddress(ctx, sdk.NewCoins(sdk.NewCoin(sdk.NativeToken, amount)), ownerAddress, address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to distribute to address")
	}

	nft.AvalableToWithdraw = nft.AvalableToWithdraw.Sub(sdk.NewCoin(sdk.NativeToken, amount))
	return new(types.MsgWithdrawalResponse), nil
}

func (k msgServer) Redelegate(goctx context.Context, request *types.MsgRedelegate) (*types.MsgRedelegateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)

	nft, ok := k.GetNFT(ctx, request.Address)
	if !ok {
		return nil, types.ErrNFTNotFound
	}

	if nft.Owner != request.Creator {
		return nil, types.ErrNFTInvalidOwner
	}

	address, _ := sdk.AccAddressFromBech32(request.Address)
	validatorSrcAddress, _ := sdk.ValAddressFromBech32(request.ValidatorSrc)
	validatorNEwAddress, _ := sdk.ValAddressFromBech32(request.ValidatorNew)

	_, err := k.stakingKeeper.BeginRedelegation(ctx, address, validatorSrcAddress, validatorNEwAddress, sdk.NewDecCoinFromCoin(request.Amount).Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to begin redelegation")
	}

	return new(types.MsgRedelegateResponse), nil
}
