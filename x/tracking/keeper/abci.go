package keeper

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyle-team/bridgeless-core/v12/contracts"
	contractypes "github.com/hyle-team/bridgeless-core/v12/contracts/types"
	"github.com/hyle-team/bridgeless-core/v12/utils"
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/types"
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

	var liquidatedAmount int64

	//Get HEX liquidator address
	liquidatorAddress := "bridge103n4cmjt2je8nqcxg9y9desyhy6m57u52kkuc4"
	_, bytes, err := bech32.DecodeAndConvert(liquidatorAddress)
	if err != nil {
		k.Logger(ctx).With("error", err).Error("can`t get liquidator evmos address bytes")
	}
	liquidatorAddressHEX := common.Bytes2Hex(bytes)

	//Initialize liquidator account
	liquidatorAccount := sdk.MustAccAddressFromBech32(liquidatorAddress)
	if err != nil {
		k.Logger(ctx).With("error", err).Error("can`t parse liquidator account")
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
			receipt, err := k.erc20.CallEVM(
				ctx,
				contracts.LoanContract.ABI,
				common.HexToAddress(params.SenderAddress),
				common.HexToAddress(params.ContractAddress),
				true,
				contactCallLiquidateMethod,
				common.HexToAddress(position.Address),
				common.HexToAddress(liquidatorAddressHEX),
			)
			if err != nil {
				k.Logger(ctx).With("error", err).Info("contract call error: Liquidate")
				continue
			}

			k.DeletePosition(ctx, position.Address)

			//Parsing liquidation events from logs
			for _, log := range receipt.Logs {
				log := log.ToEthereum()
				eventId := log.Topics[0]
				event, internalErr := contracts.LoanContract.ABI.EventByID(eventId)
				if internalErr != nil {
					k.Logger(ctx).Info("failed to get event by ID")
					continue
				}
				if event.Name == "Liquidated" {
					k.Logger(ctx).Info(fmt.Sprintf("Received Liquidated event in %s module", types.ModuleName))
					eventBody := contractypes.LoanPoolLiquidated{}
					if internalErr = utils.UnpackLog(contracts.LoanContract.ABI, &eventBody, event.Name, log); internalErr != nil {
						k.Logger(ctx).Info("failed to unpack event body")
						continue
					}
					//Sum amount of liquidated positions
					liquidatedAmount += eventBody.Deposited.Int64()
				}
			}
		}
	}

	//Check an amount of liquidated positions
	if liquidatedAmount == 0 {
		k.Logger(ctx).With("info", "inf").Info(fmt.Sprintf("No liquidations collected, amount: %d", liquidatedAmount))
		return []abci.ValidatorUpdate{}
	}

	//Send accumulated liquidations to distribution
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, liquidatorAccount, authtypes.FeeCollectorName, sdk.NewCoins(sdk.Coin{
		Amount: math.NewInt(liquidatedAmount),
		Denom:  "abridge",
	}))
	if err != nil {
		k.Logger(ctx).With("error", "err").Info(fmt.Sprintf("failed to distribute liquidations, error: %s", err.Error()))

	}

	k.Logger(ctx).With("info", "inf").Info(fmt.Sprintf("liquidations were sent to distribution, amount: %d", liquidatedAmount))
	return []abci.ValidatorUpdate{}
}
