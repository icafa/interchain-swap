package keeper

import (
	"context"

    "github.com/junkai121/interchain-swap/x/interquery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) RegisterICA(goCtx context.Context,  msg *types.MsgRegisterICA) (*types.MsgRegisterICAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgRegisterICAResponse{}, nil
}
