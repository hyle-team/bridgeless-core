package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyle-team/bridgeless-core/contracts"
	erc20keeper "github.com/hyle-team/bridgeless-core/x/erc20/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"math/big"

	"github.com/hyle-team/bridgeless-core/x/tracking/types"
)

const (
	contactCallMethod = "getLoanPositionInfo"
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

func (k *Keeper) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	params := k.GetParams(ctx)
	k.Logger(ctx).Info("contact address ", params.ContractAddress)

	for _, positionID := range params.GetPositions() {
		id, ok := big.NewInt(0).SetString(positionID, 10)
		if !ok {
			k.Logger(ctx).Error(fmt.Sprintf("invalid position id: %s", positionID))
		}

		_, err := k.erc20.CallEVM(
			ctx,
			contracts.LoanContract.ABI,
			common.HexToAddress(params.SenderAddress),
			common.HexToAddress(params.ContractAddress),
			false,
			contactCallMethod,
			id,
		)
		if err != nil {

			k.Logger(ctx).Error("contract call error", err)
			return []abci.ValidatorUpdate{}
		}
	}

	return []abci.ValidatorUpdate{}
}
