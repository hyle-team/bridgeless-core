package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/hyle-team/bridgeless-core/v12/x/tracking/types"
	"github.com/spf13/cobra"
)

func CmdQueryPositions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "positions",
		Short: "shows all active positions",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryPositions{
				Pagination: pageReq,
			}

			res, err := queryClient.GetAllPositions(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryPositionByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "positions [address]",
		Short: "shows the active position by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryPositionByAddress{
				Address: args[0],
			}
			res, err := queryClient.GetPositionByAddress(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
