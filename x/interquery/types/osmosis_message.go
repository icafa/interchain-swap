package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	TypeMsgSwapExactAmountIn = "swap_exact_amount_in"
)

var _ sdk.Msg = &OsmosisMsgSwapExactAmountIn{}

func (msg OsmosisMsgSwapExactAmountIn) Route() string { return RouterKey }
func (msg OsmosisMsgSwapExactAmountIn) Type() string  { return TypeMsgSwapExactAmountIn }
func (msg OsmosisMsgSwapExactAmountIn) ValidateBasic() error {
	return nil
}

func (msg OsmosisMsgSwapExactAmountIn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg OsmosisMsgSwapExactAmountIn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
