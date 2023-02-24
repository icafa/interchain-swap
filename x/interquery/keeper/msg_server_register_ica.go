package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/junkai121/interchain-swap/x/interquery/types"
)

func (k msgServer) RegisterICA(goCtx context.Context, msg *types.MsgRegisterICA) (*types.MsgRegisterICAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.icaControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, msg.Creator, msg.Version); err != nil {
		return nil, err
	}

	return &types.MsgRegisterICAResponse{}, nil
}
