package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v4/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	"github.com/junkai121/interchain-swap/x/interquery/types"
)

func (k msgServer) SendOsmosisSwap(goCtx context.Context, msg *types.MsgSendOsmosisSwap) (*types.MsgSendOsmosisSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	portID, err := icatypes.NewControllerPortID(msg.Sender)
	if err != nil {
		return nil, err
	}

	channelID, found := k.icaControllerKeeper.GetActiveChannelID(ctx, msg.ConnectionId, portID)
	if !found {
		return nil, icatypes.ErrActiveChannelNotFound.Wrapf("failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
	if !found {
		return nil, channeltypes.ErrChannelCapabilityNotFound.Wrap("module does not own channel capability")
	}

	addr, found := k.icaControllerKeeper.GetInterchainAccountAddress(ctx, msg.ConnectionId, portID)
	if !found {
		return nil, types.ErrICANotFound
	}

	m := &types.OsmosisMsgSwapExactAmountIn{
		Sender:            addr,
		Routes:            msg.Routes,
		TokenIn:           msg.TokenIn,
		TokenOutMinAmount: msg.TokenOutMinAmount,
	}

	data, err := icatypes.SerializeCosmosTx(k.cdc, []sdk.Msg{m})
	if err != nil {
		return nil, err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
	}

	timeoutTimestamp := ctx.BlockTime().Add(time.Minute).UnixNano()
	_, err = k.icaControllerKeeper.SendTx(ctx, chanCap, msg.ConnectionId, portID, packetData, uint64(timeoutTimestamp))
	if err != nil {
		return nil, err
	}

	return &types.MsgSendOsmosisSwapResponse{}, nil
}
