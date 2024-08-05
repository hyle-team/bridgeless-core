package cli

import (
	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"github.com/spf13/cobra"
)

func TxTransactionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "transactions",
		Short:                      "transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSubmitTx(),
	)

	return cmd
}

func CmdSubmitTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit [from_key_or_address] [json-tx]",
		Short: "Submit a transaction to the bridge module",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var tr types.Transaction
			if err = types.ModuleCdc.UnmarshalJSON([]byte(args[1]), &tr); err != nil {
				return errors.Wrap(err, "failed to unmarshal transaction")
			}

			msg := types.NewMsgSubmitTransaction(clientCtx.GetFromAddress().String(), tr)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
