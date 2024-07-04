package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	globaltypes "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	accountKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/evmos/evmos/v12/x/accumulator/types"
	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/net/context"
	"time"
)

type (
	Keeper interface {
		Logger(c sdk.Context) log.Logger
		GetModuleInfo(ctx sdk.Context, moduleName string) *types.ModuleInfo
		GetParams(c sdk.Context) types.Params
		SetParams(c sdk.Context, params types.Params)
		Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)

		SendFromModuleToAddress(ctx sdk.Context, moduleNameFrom string, address sdk.AccAddress, amount sdk.Coins) error
		SendToModuleAddress(ctx sdk.Context, moduleNameFrom, moduleNameTo string, amount sdk.Coins) error
		SendFromAddressToAddress(ctx sdk.Context, moduleNameFrom, address string, amount sdk.Coins) error
		ValidateBalance(ctx sdk.Context, moduleName string, amount sdk.Coins) error
	}

	BaseKeeper struct {
		cdc             codec.BinaryCodec
		storeKey        storetypes.StoreKey
		memKey          storetypes.StoreKey
		bankKeeper      bankkeeper.Keeper
		ak              accountKeeper.AccountKeeper
		lastVestingTime time.Time
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ak accountKeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
) *BaseKeeper {
	return &BaseKeeper{
		bankKeeper: bankKeeper,
		cdc:        cdc,
		storeKey:   storeKey,
		ak:         ak,
		memKey:     memKey,
	}
}

func (k BaseKeeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k BaseKeeper) GetModuleInfo(ctx sdk.Context, moduleName string) *types.ModuleInfo {
	return k.GetParams(ctx).ModulesInfo[moduleName]
}

func (k BaseKeeper) ValidateBalance(ctx sdk.Context, moduleName string, amount sdk.Coins) error {
	k.Logger(ctx).Info("ValidateBalance", "module", moduleName, "amount", amount)
	moduleInfo := k.GetModuleInfo(ctx, moduleName)
	if moduleInfo == nil {
		return fmt.Errorf("module info not found")
	}

	address, err := sdk.AccAddressFromBech32(moduleInfo.Address)

	if err != nil {
		return errors.Wrap(err, "invalid address")
	}

	if moduleInfo.Address == "" {
		moduleAddress := k.ak.GetModuleAccount(ctx, moduleName)
		address = moduleAddress.GetAddress()
	}

	balance := k.bankKeeper.GetBalance(ctx, address, globaltypes.NativeToken)
	fmt.Println("balance: ", balance)
	fmt.Println("amount: ", amount)
	if !balance.IsValid() || !balance.IsPositive() || amount.AmountOf(globaltypes.NativeToken).Sub(balance.Amount).IsPositive() {
		return fmt.Errorf("invalid balance")
	}

	return nil
}
