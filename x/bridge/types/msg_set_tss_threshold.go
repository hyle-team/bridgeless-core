package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetTssThreshold = "set_tss_threshold"

var _ sdk.Msg = &MsgSetTssThreshold{}

func NewMsgSetTssThreshold(creator string, amount string) *MsgSetTssThreshold {
	return &MsgSetTssThreshold{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgSetTssThreshold) Route() string {
	return RouterKey
}

func (msg *MsgSetTssThreshold) Type() string {
	return TypeMsgSetParties
}

func (msg *MsgSetTssThreshold) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetTssThreshold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetTssThreshold) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid amount (%s)", err)
	}

	if !amount.GT(math.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "threshold amount must be greater than zero")
	}

	return nil
}
