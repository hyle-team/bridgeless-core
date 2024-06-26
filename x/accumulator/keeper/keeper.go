package keeper

import (
	"fmt"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	accountKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	globaltypes "github.com/evmos/evmos/v18/types"
	"github.com/evmos/evmos/v18/x/accumulator/types"
	"golang.org/x/net/context"
)

type (
	Keeper interface {
		Logger(c sdk.Context) log.Logger
		GetModuleInfo(ctx sdk.Context, moduleName string) *types.ModuleInfo
		GetParams(c sdk.Context) types.Params
		SetParams(c sdk.Context, params types.Params)
		Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error)
		MintTokens(ctx sdk.Context, amount int64, moduleName string) error
		//DistributeTokens(ctx sdk.Context, amount int64, moduleNameFrom, moduleNameTo string, address *string) error

		sendToModuleAddress(ctx sdk.Context, moduleNameFrom, moduleNameTo string, amount int64) error
		sendToAddress(ctx sdk.Context, moduleNameFrom, address string, amount int64) error
		validateBalance(ctx sdk.Context, moduleName string, amount int64) error
	}

	BaseKeeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		bankKeeper bankkeeper.Keeper
		ak         accountKeeper.AccountKeeper
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

func (k BaseKeeper) validateBalance(ctx sdk.Context, moduleName string, amount int64) error {
	moduleInfo := k.GetModuleInfo(ctx, moduleName)
	address := sdk.AccAddress(moduleInfo.Address)

	if moduleInfo.Address == "" {
		moduleAddress := k.ak.GetModuleAccount(ctx, moduleName)
		address = moduleAddress.GetAddress()
	}

	balance := k.bankKeeper.GetBalance(ctx, address, globaltypes.NativeToken)
	if !balance.IsValid() || !balance.IsPositive() || balance.Amount.Int64() < amount {
		return fmt.Errorf("invalid balance")
	}
	return nil
}
