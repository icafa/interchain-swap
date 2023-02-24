package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/junkai121/interchain-swap/x/interquery/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSendQueryOsmosisPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-query-all-balances [channel-id] [pool-id] [base-asset-denom] [quote-asset-denom] [address]",
		Short: "Query the balances of an account on the remote chain via ICQ",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolId, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgSendQueryOsmosisPrice(
				clientCtx.GetFromAddress().String(),
				args[0], // channel id
				uint64(poolId),
				args[2],
				args[3],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "send query all balances")

	return cmd
}
