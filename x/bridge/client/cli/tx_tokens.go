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
	"strconv"
)

func CmdSetToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-token [from_key_or_address] [token-file] ",
		Short: "Set a new token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			token := &types.Token{}

			tokenFile, err := os.Open(args[1])
			if err != nil {
				return err
			}
			defer tokenFile.Close()

			rawToken, err := io.ReadAll(tokenFile)
			if err != nil {
				return err
			}

			err = json.Unmarshal(rawToken, token)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetToken(
				clientCtx.GetFromAddress().String(),
				token,
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

func CmdRemoveToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-token [from_key_or_address] [token_id]",
		Short: "Remove the token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[1], 10, 64)

			msg := types.NewMsgRemoveToken(
				clientCtx.GetFromAddress().String(),
				id,
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
