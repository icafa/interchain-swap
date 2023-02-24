package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendOsmosisSwap = "send_osmosis_swap"

var _ sdk.Msg = &MsgSendOsmosisSwap{}

func NewMsgSendOsmosisSwap(creator string, title string, body string) *MsgSendOsmosisSwap {
	return &MsgSendOsmosisSwap{
		Sender: creator,
	}
}

func (msg *MsgSendOsmosisSwap) Route() string {
	return RouterKey
}

func (msg *MsgSendOsmosisSwap) Type() string {
	return TypeMsgSendOsmosisSwap
}

func (msg *MsgSendOsmosisSwap) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendOsmosisSwap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendOsmosisSwap) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
