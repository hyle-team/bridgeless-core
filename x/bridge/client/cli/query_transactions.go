package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/hyle-team/bridgeless-core/v12/x/bridge/types"
	"github.com/spf13/cobra"
)

func CmdQueryTransactions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactions",
		Short: "Query bridge transactions",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Transactions(cmd.Context(), &types.QueryTransactionsRequest{Pagination: pageReq})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "txs")

	return cmd
}

func CmdQueryTransactionById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transaction [id]",
		Short: "Query bridge transaction by its id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.TransactionById(cmd.Context(), &types.QueryTransactionByIdRequest{Id: args[0]})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
