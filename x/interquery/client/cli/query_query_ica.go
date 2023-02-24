package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/junkai121/interchain-swap/x/interquery/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQueryIca() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-ica [connection-id] [owner]",
		Short: "Query queryIca",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryQueryIcaRequest{
				ConnectionId: args[0],
				Owner:        args[1],
			}

			res, err := queryClient.QueryIca(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
