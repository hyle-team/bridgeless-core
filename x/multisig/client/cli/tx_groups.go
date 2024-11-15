package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hyle-team/bridgeless-core/v12/x/multisig/types"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

func TxGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "groups",
		Short:                      "Group transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdCreateGroup(),
		CmdUpdateGroup(),
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdCreateGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [from_creator_address] [members_addresses] [threshold]",
		Short: "Create a group",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			arg2 := args[1]
			members := strings.Split(arg2, ",")

			for _, member := range members {
				_, err = sdk.AccAddressFromBech32(member)
				if err != nil {
					return err
				}
			}

			arg3, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}
			threshold := uint64(arg3)

			msg := types.NewMsgCreateGroup(clientCtx.GetFromAddress().String(), members, threshold)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdUpdateGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [creator_address] [group_address] [members_addresses] [threshold]",
		Short: "Update a group",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Flags().Set(flags.FlagFrom, args[0])
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			arg2, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			groupAddress := arg2.String()

			arg3 := args[2]
			members := strings.Split(arg3, ",")

			for _, member := range members {
				_, err = sdk.AccAddressFromBech32(member)
				if err != nil {
					return err
				}
			}

			arg4, err := strconv.Atoi(args[3])
			if err != nil {
				return err
			}
			threshold := uint64(arg4)

			msg := types.NewMsgChangeGroup(clientCtx.GetFromAddress().String(), groupAddress, members, threshold)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
