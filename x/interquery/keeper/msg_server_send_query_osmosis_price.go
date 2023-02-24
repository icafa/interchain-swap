package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	"github.com/junkai121/interchain-swap/x/interquery/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

func (k msgServer) SendQueryOsmosisPrice(goCtx context.Context, msg *types.MsgSendQueryOsmosisPrice) (*types.MsgSendQueryOsmosisPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(k.GetPort(ctx), msg.ChannelId))
	if !found {
		return nil, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	q := types.QuerySpotPriceRequest{
		PoolId:          msg.PoolId,
		BaseAssetDenom:  msg.BaseAssetDenom,
		QuoteAssetDenom: msg.QuoteAssetDenom,
	}
	reqs := []abcitypes.RequestQuery{
		{
			Path: "/osmosis.gamm.v1beta1.Query/SpotPrice",
			Data: k.cdc.MustMarshal(&q),
		},
	}

	// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
	// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
	timeoutTimestamp := ctx.BlockTime().Add(time.Minute).UnixNano()
	seq, err := k.SendQuery(ctx, types.PortID, msg.ChannelId, chanCap, reqs, clienttypes.ZeroHeight(), uint64(timeoutTimestamp))
	if err != nil {
		return nil, err
	}

	k.SetQueryRequest(ctx, seq, q)

	return &types.MsgSendQueryOsmosisPriceResponse{
		Sequence: seq,
	}, nil
}
