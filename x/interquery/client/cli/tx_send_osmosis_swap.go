package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/junkai121/interchain-swap/x/interquery/types"
	"github.com/spf13/cobra"
)

const (
	// Will be parsed to uint64.
	FlagSwapRoutePoolIds = "swap-route-pool-ids"
	// Will be parsed to []string.
	FlagSwapRouteDenoms = "swap-route-denoms"
)

func FlagSetQuerySwapRoutes() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagSwapRoutePoolIds, "", "swap route pool id")
	fs.String(FlagSwapRouteDenoms, "", "swap route amount")
	return fs
}

var _ = strconv.Itoa(0)

func CmdSendOsmosisSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-exact-amount-in [connection-id] [token-in] [token-out-min-amount]",
		Short: "Broadcast message sendOsmosisSwap",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)
			txf, msg, err := NewBuildSwapExactAmountInMsg(clientCtx, args[0], args[1], args[2], txf, cmd.Flags())
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().AddFlagSet(FlagSetQuerySwapRoutes())
	_ = cmd.MarkFlagRequired(FlagSwapRoutePoolIds)
	_ = cmd.MarkFlagRequired(FlagSwapRouteDenoms)

	return cmd
}

func NewBuildSwapExactAmountInMsg(clientCtx client.Context, connectionId, tokenInStr, tokenOutMinAmtStr string, txf tx.Factory, fs *flag.FlagSet) (tx.Factory, sdk.Msg, error) {
	routes, err := swapAmountInRoutes(fs)
	if err != nil {
		return txf, nil, err
	}

	tokenIn, err := sdk.ParseCoinNormalized(tokenInStr)
	if err != nil {
		return txf, nil, err
	}

	tokenOutMinAmt, ok := sdk.NewIntFromString(tokenOutMinAmtStr)
	if !ok {
		return txf, nil, fmt.Errorf("invalid token out min amount, %s", tokenOutMinAmtStr)
	}
	msg := &types.MsgSendOsmosisSwap{
		Sender:            clientCtx.GetFromAddress().String(),
		ConnectionId:      connectionId,
		Routes:            routes,
		TokenIn:           tokenIn,
		TokenOutMinAmount: tokenOutMinAmt,
	}

	return txf, msg, nil
}

func swapAmountInRoutes(fs *flag.FlagSet) ([]types.SwapAmountInRoute, error) {
	swapRoutePoolIds, err := fs.GetString(FlagSwapRoutePoolIds)
	swapRoutePoolIdsArray := strings.Split(swapRoutePoolIds, ",")
	if err != nil {
		return nil, err
	}

	swapRouteDenoms, err := fs.GetString(FlagSwapRouteDenoms)
	swapRouteDenomsArray := strings.Split(swapRouteDenoms, ",")
	if err != nil {
		return nil, err
	}

	if len(swapRoutePoolIdsArray) != len(swapRouteDenomsArray) {
		return nil, errors.New("swap route pool ids and denoms mismatch")
	}

	routes := []types.SwapAmountInRoute{}
	for index, poolIDStr := range swapRoutePoolIdsArray {
		pID, err := strconv.Atoi(poolIDStr)
		if err != nil {
			return nil, err
		}
		routes = append(routes, types.SwapAmountInRoute{
			PoolId:        uint64(pID),
			TokenOutDenom: swapRouteDenomsArray[index],
		})
	}
	return routes, nil
}
