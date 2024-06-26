package keeper

import (
	"cosmossdk.io/errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	accountKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	globaltypes "github.com/evmos/evmos/v18/types"
	"github.com/evmos/evmos/v18/utils"

	"github.com/evmos/evmos/v18/x/accumulator/types"
	"github.com/tendermint/tendermint/libs/log"
	"golang.org/x/net/context"
)

type (
	Keeper interface {
		Logger(c sdk.Context) log.Logger
		GetModuleInfo(moduleName string) types.ModuleInfo
		GetParams(c sdk.Context) types.Params
		SetParams(c sdk.Context, params types.Params)
		Params(ctx context.Context, req *types.QueryParamsRequest)
		MintTokens(ctx sdk.Context, amount int64, moduleName string) error
		DistributeTokens(ctx sdk.Context, amount int64, moduleNameFrom, moduleNameTo string, address *string) error

		sendToModuleAddress(c sdk.Context, moduleFrom, moduleTo string, coins sdk.Coins) error
		sendToAddress(c sdk.Context, moduleFrom string, address sdk.AccAddress, coins sdk.Coins) error
		validateBalance(ctx sdk.Context, moduleName string, amount int64) error
	}

	BaseKeeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		mintKeeper mintkeeper.Keeper
		bankKeeper bankkeeper.Keeper
		ak         accountKeeper.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	mintKeeper mintkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
) *BaseKeeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &BaseKeeper{
		bankKeeper: bankKeeper,
		mintKeeper: mintKeeper,
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k BaseKeeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k BaseKeeper) GetModuleInfo(ctx sdk.Context, moduleName string) *types.ModuleInfo {
	return k.GetParams(ctx).ModulesInfo[moduleName]
}

func (k BaseKeeper) MintTokens(ctx sdk.Context, amount int64, moduleName string) error {
	if err := k.bankKeeper.MintCoins(ctx, moduleName, utils.NewnNativeTokens(amount)); err != nil {
		return errors.Wrap(err, "mint native tokens")
	}

	return nil
}

func (k BaseKeeper) validateBalance(ctx sdk.Context, moduleName string, amount int64) error {
	moduleInfo := k.GetParams(ctx).ModulesInfo[moduleName]

	balance := k.bankKeeper.GetBalance(ctx, sdk.AccAddress(moduleInfo.Address), globaltypes.NativeToken)
	if !balance.IsValid() || !balance.IsPositive() || balance.Amount.Int64() < amount {
		return fmt.Errorf("invalid balance")
	}
	return nil
}
