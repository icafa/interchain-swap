package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendQueryOsmosisPrice = "send_query_all_balances"

var _ sdk.Msg = &MsgSendQueryOsmosisPrice{}

func NewMsgSendQueryOsmosisPrice(creator string, channelId string, poolId uint64, baseAssetDenom, quoteAssetDenom string) *MsgSendQueryOsmosisPrice {
	return &MsgSendQueryOsmosisPrice{
		Creator:         creator,
		ChannelId:       channelId,
		PoolId:          poolId,
		BaseAssetDenom:  baseAssetDenom,
		QuoteAssetDenom: quoteAssetDenom,
	}
}

func (msg *MsgSendQueryOsmosisPrice) Route() string {
	return RouterKey
}

func (msg *MsgSendQueryOsmosisPrice) Type() string {
	return TypeMsgSendQueryOsmosisPrice
}

func (msg *MsgSendQueryOsmosisPrice) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendQueryOsmosisPrice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendQueryOsmosisPrice) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
