package keeper

import (
	"bytes"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/hyle-team/bridgeless-core/contracts"
	contractypes "github.com/hyle-team/bridgeless-core/contracts/types"
	"github.com/hyle-team/bridgeless-core/utils"
	"github.com/hyle-team/bridgeless-core/x/tracking/types"
)

// PostTxProcessing is used to listen EVM smart contract events,
// filter and process `PositionCreated` events emitted by configured in module params contract address.
// Will be called by EVM module as hook.
func (k Keeper) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	params := k.GetParams(ctx)

	k.Logger(ctx).Info("PostTxProcessing", "msg", msg, "receipt", receipt)

	contractAddress, err := hexutil.Decode(params.ContractAddress)
	if err != nil {
		// If return an error here, the whole EVM module won't work
		k.Logger(ctx).Error("failed to decode contract address")
		return nil
	}

	// https://docs.evmos.org/protocol/modules/evm#posttxprocessing

	if len(receipt.Logs) == 0 {
		k.Logger(ctx).Error("logs is empty")
	}

	// This approach is used because the contract address we use for validation and the event we
	//want to catch are in different logs.
	isAppropriateAddress := false
	for _, log := range receipt.Logs {
		// Validating message receiver address (should be our state smart contract)
		if log != nil && log.Address.Bytes() != nil && bytes.Compare(log.Address.Bytes(), contractAddress) == 0 {
			isAppropriateAddress = true
		}
	}
	if !isAppropriateAddress {
		return nil
	}

	for _, log := range receipt.Logs {
		eventId := log.Topics[0]
		event, internalErr := contracts.LoanContract.ABI.EventByID(eventId)
		if internalErr != nil {
			k.Logger(ctx).Error("failed to get event by ID")
			continue
		}

		if event.Name != params.EventName {
			k.Logger(ctx).Info(fmt.Sprintf("unmatched event: got %s, expected %s", event.Name, params.EventName))
			continue
		}

		eventBody := contractypes.ILoanPoolDeposited{}
		if internalErr = utils.UnpackLog(contracts.LoanContract.ABI, &eventBody, event.Name, log); internalErr != nil {
			k.Logger(ctx).Error("failed to unpack event body")
			continue
		}

		k.SetPosition(ctx, eventBody.User.Hex(), types.Position{
			Address:        eventBody.User.Hex(),
			Amount:         eventBody.Amount.Int64(), // TODO maybe need to convert to string or update the store to save big.Int
			LastTimeUpdate: ctx.BlockTime().Unix(),
		})

		k.Logger(ctx).Info(fmt.Sprintf("Received PostTxProcessing event in %s module: %v", types.ModuleName, eventBody))
		k.SetParams(ctx, params)
	}

	return nil
}
