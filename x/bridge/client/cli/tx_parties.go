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
		Use:   "set [list-type (default,goodbye,newbies,blacklist)] [from_key_or_address] [pas-list]",
		Short: "Set a new parties list",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[1])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return errors.Wrap(err, "cannot get client tx context")
			}

			var partiesList []*types.Party
			if len(args) == 3 {
				parties := strings.Split(args[2], ",")
				for _, party := range parties {
					_, err = sdk.AccAddressFromBech32(party)
					if err != nil {
						return errors.Wrap(err, "failed to set parties")
					}
				}

				for _, party := range parties {
					partiesList = append(partiesList, &types.Party{
						Address: party,
					})
				}
			}

			var msg sdk.Msg
			switch args[0] {
			case "default":
				msg = types.NewMsgSetPartiesList(clientCtx.GetFromAddress().String(), partiesList)
			case "newbies":
				msg = types.NewMsgSetNewbiesList(clientCtx.GetFromAddress().String(), partiesList)
			case "goodbye":
				msg = types.NewMsgSetGoodbyeList(clientCtx.GetFromAddress().String(), partiesList)
			case "blacklist":
				msg = types.NewMsgSetBlacklistPartiesList(clientCtx.GetFromAddress().String(), partiesList)
			default:
				return errors.New(fmt.Sprintf("invalid parties list type: %s", args[0]))
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		}}
	return cmd
}
