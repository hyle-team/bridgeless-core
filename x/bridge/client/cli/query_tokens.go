package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/hyle-team/bridgeless-core/x/bridge/types"
	"github.com/spf13/cobra"
)

func CmdQueryGetTokenById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token [id]",
		Short: "shows the token info by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryGetTokenById{
				Id: args[0],
			}
			res, err := queryClient.GetTokenById(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens --flags",
		Short: "Query the all tokens",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			req := &types.QueryGetTokens{
				Pagination: pageReq,
			}

			res, err := queryClient.GetTokens(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryTokenPairs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-pairs [src_chain] [src_address] --flags",
		Short: "Query the all token pairs",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.GetAllTokenPairs(context.Background(), &types.QueryGetTokenPairs{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdQueryTokenPair() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-pairs [src_chain] [src_address] [dst_chain] --flags",
		Short: "Query the token pair",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryGetTokenPair{
				SrcChain:   args[0],
				SrcAddress: args[1],
				DstChain:   args[2],
			}

			res, err := queryClient.GetTokenPairForDstChain(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
