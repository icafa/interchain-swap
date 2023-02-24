package keeper_test

import (
	"testing"

	testkeeper "github.com/junkai121/interchain-swap/testutil/keeper"
	"github.com/junkai121/interchain-swap/x/interquery/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.InterqueryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
