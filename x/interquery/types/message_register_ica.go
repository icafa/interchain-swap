package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterICA = "register_ica"

var _ sdk.Msg = &MsgRegisterICA{}

func NewMsgRegisterICA(creator string, title string, body string) *MsgRegisterICA {
	return &MsgRegisterICA{
		Creator: creator,
	}
}

func (msg *MsgRegisterICA) Route() string {
	return RouterKey
}

func (msg *MsgRegisterICA) Type() string {
	return TypeMsgRegisterICA
}

func (msg *MsgRegisterICA) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterICA) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterICA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
