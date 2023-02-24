package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/junkai121/interchain-swap/x/interquery/types"
)

func (k msgServer) SendOsmosisSwap(goCtx context.Context, msg *types.MsgSendOsmosisSwap) (*types.MsgSendOsmosisSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendOsmosisSwapResponse{}, nil
}
