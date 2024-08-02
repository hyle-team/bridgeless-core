package cli

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func CmdSetChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-chain [from_key_or_address] [token_json_file]",
		Short: "Set a new chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chain := &types.Chain{}
			rawChain, err := os.Open(args[1])
			if err != nil {
				return err
			}
			defer rawChain.Close()

			byteValue, err := io.ReadAll(rawChain)
			if err != nil {
				return err
			}

			err = json.Unmarshal(byteValue, chain)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetChain(
				clientCtx.GetFromAddress().String(),
				chain,
			)

			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdRemoveChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-chain [from_key_or_address] [chain_id]",
		Short: "Remove the chain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveChain(
				clientCtx.GetFromAddress().String(),
				args[1],
			)

			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
