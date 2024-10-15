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
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	// Gas costs are handled within msg handler so costs should be ignored

	fmt.Println("contrat address ", contracts.LoanContractAddress)

	res, err := k.erc20.CallEVM(
		ctx,
		contracts.LoanContract.ABI,
		common.HexToAddress("0xa87044815A445E5C632c51793bd1065C35DC405d"),
		contracts.LoanContractAddress,
		true,
		"decrement",
		big.NewInt(34),
	)
	if err != nil {
		fmt.Println("ERRRRRROOOOOOOOORRRRRR ABIIIIII")
		fmt.Println(err)
		return []abci.ValidatorUpdate{}
	}
	fmt.Println("abi RESSSSSSSSS: ", res.Hash, res.Logs, res.GasUsed, res.String())

	return []abci.ValidatorUpdate{}
}
