package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyle-team/bridgeless-core/v12/contracts"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (k *Keeper) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	params := k.GetParams(ctx)
	k.Logger(ctx).Info("contact address ", params.ContractAddress)
	positions, _, err := k.GetPositionsWithPagination(ctx, &query.PageRequest{Limit: query.MaxLimit})
	if err != nil {
		k.Logger(ctx).Error("failed to get positions", err)
		return []abci.ValidatorUpdate{}
	}

	_, err = k.erc20.CallEVM(
		ctx,
		contracts.LoanContract.ABI,
		common.HexToAddress(params.SenderAddress),
		common.HexToAddress(params.ContractAddress),
		true,
		contactCallOracleMethod,
		params.Oracle,
	)

	for _, position := range positions {
		if position.LastTimeUpdate-ctx.BlockHeight() >= int64(params.Threshold) {
			txResponse, err := k.erc20.CallEVM(
				ctx,
				contracts.LoanContract.ABI,
				common.HexToAddress(params.SenderAddress),
				common.HexToAddress(params.ContractAddress),
				false,
				contactCallIsUserCanBeLiquidatedMethod,
				position.Address,
			)
			if err != nil {
				k.Logger(ctx).Error("contract call error", err)
				continue
			}

			unpackedRes, err := contracts.LoanContract.ABI.Unpack("isUserCanBeLiquidated", txResponse.Ret)
			if err != nil {
				k.Logger(ctx).Error("unpack contract", err)
				continue
			}

			fmt.Println(txResponse)
			// TODO parse tx response and replace true with the result
			if unpackedRes[0].(bool) {
				_, err = k.erc20.CallEVM(
					ctx,
					contracts.LoanContract.ABI,
					common.HexToAddress(params.SenderAddress),
					common.HexToAddress(params.ContractAddress),
					true,
					contactCallLiquidateMethod,
					position.Address,
				)
				if err != nil {
					k.Logger(ctx).Error("contract call error", err)
					continue
				}

				k.DeletePosition(ctx, position.Address)
			}
		}
	}

	if params.LastUpdateBunch*params.BunchSize >= int32(len(positions)) {
		// reset queue if all positions are updated
		params.LastUpdateBunch = 0
		k.SetParams(ctx, params)
	}

	var bunchLastIndex = params.LastUpdateBunch*params.BunchSize + params.BunchSize
	if int32(len(positions))-params.LastUpdateBunch*params.BunchSize < params.BunchSize {
		bunchLastIndex = int32(len(positions)) - params.LastUpdateBunch*params.BunchSize
	}

	positionsToUpdate := positions[params.LastUpdateBunch*params.BunchSize : bunchLastIndex]
	for _, position := range positionsToUpdate {
		_, err = k.erc20.CallEVM(
			ctx,
			contracts.LoanContract.ABI,
			common.HexToAddress(params.SenderAddress),
			common.HexToAddress(params.ContractAddress),
			true,
			contactCallUpdatePositionMethod,
			position.Address,
		)
		if err != nil {
			k.Logger(ctx).Error("contract call error", err)
			continue
		}
	}

	return []abci.ValidatorUpdate{}
}
