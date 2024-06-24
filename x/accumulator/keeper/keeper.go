package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/evmos/evmos/v18/x/accumulator/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc         codec.BinaryCodec
		storeKey    storetypes.StoreKey
		memKey      storetypes.StoreKey
		paramstore  paramtypes.Subspace
		mintKeeper  mintkeeper.Keeper
		bankKeeper  bankkeeper.Keeper
		modulesInfo map[string]ModuleInfo
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	mintKeeper mintkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
	modulesInfo map[string]ModuleInfo,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		bankKeeper:  bankKeeper,
		mintKeeper:  mintKeeper,
		cdc:         cdc,
		storeKey:    storeKey,
		memKey:      memKey,
		paramstore:  ps,
		modulesInfo: modulesInfo,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetModuleInfo(moduleName string) ModuleInfo {
	return k.modulesInfo[moduleName]
}

func (k Keeper) SetModuleInfo(modulesInfo map[string]ModuleInfo) {
	k.modulesInfo = modulesInfo
}
