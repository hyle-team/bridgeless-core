package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	erc20keeper "github.com/hyle-team/bridgeless-core/v12/x/erc20/keeper"
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/types"
	"github.com/tendermint/tendermint/libs/log"
)

const (
	contactCallUpdatePositionMethod        = "updateUserPosition"
	contactCallIsUserCanBeLiquidatedMethod = "canUserBeLiquidated"
	contactCallOracleMethod                = "oracle"
	contactCallLiquidateMethod             = "liquidate"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace
		erc20      erc20keeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	erc20 erc20keeper.Keeper,

) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		erc20:    erc20,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
