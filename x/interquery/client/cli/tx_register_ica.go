package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/junkai121/interchain-swap/x/interquery/types"
)

var _ = strconv.Itoa(0)

func CmdRegisterICA() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-ica [title] [body]",
		Short: "Broadcast message registerICA",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argTitle := args[0]
             argBody := args[1]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterICA(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argBody,
				
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}