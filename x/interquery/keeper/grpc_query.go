package keeper

import (
	"github.com/junkai121/interchain-swap/x/interquery/types"
)

var _ types.QueryServer = Keeper{}
