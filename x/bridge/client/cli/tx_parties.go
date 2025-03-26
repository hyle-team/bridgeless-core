package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strings"
)

func TxPartiesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "parties",
		Short:                      "Parties transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSubmitParties())
	// this line is used by starport scaffolding # 1
	return cmd
}

func CmdSubmitParties() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [list-type (default,goodbye,newbies,blacklist)] [from_key_or_address] [parties-list]",
		Short: "Set a new parties list",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[1])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return errors.Wrap(err, "cannot get client tx context")
			}

			arg2 := args[2]
			parties := strings.Split(arg2, ",")
			for _, party := range parties {
				_, err = sdk.AccAddressFromBech32(party)
				if err != nil {
					return errors.Wrap(err, "failed to set parties")
				}
			}

			var partiesList []*types.Party

			for _, party := range parties {
				partiesList = append(partiesList, &types.Party{
					Address: party,
				})
			}
			switch args[0] {
			case "default":
				msg := types.NewMsgSetParties(clientCtx.GetFromAddress().String(), partiesList)
				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			case "newbies":
				msg := types.NewMsgSetNewbies(clientCtx.GetFromAddress().String(), partiesList)
				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			case "goodbye":
				msg := types.NewMsgSetGoodbye(clientCtx.GetFromAddress().String(), partiesList)
				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			case "blacklist":
				msg := types.NewMsgSetBlacklistParties(clientCtx.GetFromAddress().String(), partiesList)
				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			default:
				return errors.New(fmt.Sprintf("invalid parties list type: %s", args[0]))
			}
		}}
	return cmd
}
