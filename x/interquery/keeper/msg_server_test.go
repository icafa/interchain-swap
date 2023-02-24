package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/junkai121/interchain-swap/testutil/keeper"
	"github.com/junkai121/interchain-swap/x/interquery/keeper"
	"github.com/junkai121/interchain-swap/x/interquery/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.InterqueryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
