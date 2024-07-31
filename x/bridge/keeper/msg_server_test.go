package keeper_test

import (
	"context"
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/hyle-team/bridgeless-core/testutil/keeper"
	"github.com/hyle-team/bridgeless-core/x/bridge/keeper"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BridgeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func Test_msg(t *testing.T) {
	tx := types.MsgSubmitTransaction{
		Submitter: "bridge1drrftpddpg8dre7spr838kks832y3v5zczqxnp",
		Transaction: types.Transaction{
			DepositChainId:    "1",
			DepositTxHash:     "qr",
			DepositBlock:      123,
			DepositToken:      "qr",
			Amount:            "10",
			Depositor:         "bridge1drrftpddpg8dre7spr838kks832y3v5zczqxnp",
			WithdrawalChainId: "1",
			WithdrawalTxHash:  "sf",
			WithdrawalToken:   "sd",
		},
	}

	k, _ := json.Marshal(tx)
	panic(string(k))
}
