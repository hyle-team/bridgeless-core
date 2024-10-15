package keeper

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyle-team/bridgeless-core/contracts"
	erc20keeper "github.com/hyle-team/bridgeless-core/x/erc20/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/hyle-team/bridgeless-core/x/tracking/types"
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
	ps paramtypes.Subspace,
	erc20 erc20keeper.Keeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		erc20:      erc20,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	// Gas costs are handled within msg handler so costs should be ignored
	params := k.GetParams(ctx)
	fmt.Println("contrat address ", params.ContractAddress)

	for _, positionID := range params.GetPositions() {
		id, ok := big.NewInt(0).SetString(positionID, 10)
		if !ok {
			ctx.Logger().Error(fmt.Sprintf("invalid position id: %s", positionID))
		}
		res, err := k.erc20.CallEVM(
			ctx,
			contracts.LoanContract.ABI,
			common.HexToAddress(params.SenderAddress),
			common.HexToAddress(params.ContractAddress),
			true,
			"updatePosition",
			id,
		)
		if err != nil {
			fmt.Println("contract call error", err)
			continue
		}
		fmt.Println("contract call response ", res.Hash, res.Logs, res.GasUsed, res.String())

	}

	return []abci.ValidatorUpdate{}
}
