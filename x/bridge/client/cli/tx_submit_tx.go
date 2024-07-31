package cli

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"github.com/spf13/cobra"
)

func CmdSubmitTx() *cobra.Command {
	return &cobra.Command{
		Use:   "submit-tx [json-tx]",
		Short: "Submit a transaction to the bridge module",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			stx, err := parseTx(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitTransaction(stx.Submitter, stx.Transaction)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

func parseTx(raw string) (types.MsgSubmitTransaction, error) {
	var stx types.MsgSubmitTransaction
	if err := types.ModuleCdc.UnmarshalJSON([]byte(raw), &stx); err != nil {
		return stx, errors.Wrap(err, "failed to unmarshal transaction")
	}

	return stx, nil
}
