package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyle-team/bridgeless-core/v12/contracts"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (k *Keeper) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	params := k.GetParams(ctx)
	positions, _, err := k.GetPositionsWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		k.Logger(ctx).With("error", err).Error("failed to get positions", err)
		return []abci.ValidatorUpdate{}
	}

	//Update oracle price by calling the oracle contract
	_, err = k.erc20.CallEVM(
		ctx,
		contracts.OracleContract.ABI,
		common.HexToAddress(params.SenderAddress),
		common.HexToAddress(params.Oracle),
		true,
		contactCallOracleMethod,
	)
	if err != nil {
		k.Logger(ctx).With("error", err).Error("contract call error: Update oracle price")
		return []abci.ValidatorUpdate{}
	}

	// Iterate over all positions and check if the user can be liquidated
	for _, position := range positions {
		txResponse, err := k.erc20.CallEVM(
			ctx,
			contracts.LoanContract.ABI,
			common.HexToAddress(params.SenderAddress),
			common.HexToAddress(params.ContractAddress),
			false,
			contactCallIsUserCanBeLiquidatedMethod,
			common.HexToAddress(position.Address),
		)
		if err != nil {
			k.Logger(ctx).With("error", err).Error("contract call error: canUserBeLiquidated")
			continue
		}

		// Decode response from the contract call
		unpackedRes, err := contracts.LoanContract.ABI.Unpack("canUserBeLiquidated", txResponse.Ret)
		if err != nil {
			k.Logger(ctx).With("error", err).Error("unpack call response")
			continue
		}

		// Convert the first return value to bool
		// if true, call liquidate method and delete the position from the store
		if len(unpackedRes) > 0 && unpackedRes[0].(bool) {
			_, err = k.erc20.CallEVM(
				ctx,
				contracts.LoanContract.ABI,
				common.HexToAddress(params.SenderAddress),
				common.HexToAddress(params.ContractAddress),
				true,
				contactCallLiquidateMethod,
				common.HexToAddress(position.Address),
				common.HexToAddress(position.Recipient),
			)
			if err != nil {
				k.Logger(ctx).With("error", err).Info("contract call error: Liquidate")
				continue
			}

			k.DeletePosition(ctx, position.Address)
		}
	}

	return []abci.ValidatorUpdate{}
}
