package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strconv"
)

func TxThresholdsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "threshold",
		Short:                      "Thresholds subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSubmitStakeThreshold(),
		CmdSubmitTssThreshold())
	// this line is used by starport scaffolding # 1
	return cmd
}

func CmdSubmitStakeThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-stake-threshold [from_key_or_address] [amount]",
		Short: "Set a stake threshold",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return errors.Wrap(err, "cannot get client tx context")
			}

			amount, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return errors.Wrap(err, "invalid stake threshold")
			}
			if amount.IsNegative() || amount.IsZero() {
				return errors.New("amount must be non-zero positive")
			}

			msg := types.NewMsgSetStakeThreshold(clientCtx.GetFromAddress().String(), args[1])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		}}
	return cmd
}

func CmdSubmitTssThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-tss-threshold [from_key_or_address] [amount]",
		Short: "Set a tss threshold",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return errors.Wrap(err, "cannot get client tx context")
			}

			amount, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil {
				return errors.Wrap(err, "invalid tss threshold")
			}

			msg := types.NewMsgSetTssThreshold(clientCtx.GetFromAddress().String(), uint32(amount))
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		}}
	return cmd
}
