package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSetPair = "set-pair"
)

var _ sdk.Msg = &MsgSetPair{}

func NewMsgSetPair(creator string, pair TokenInfo) *MsgSetPair {
	return &MsgSetPair{
		Creator: creator,
		Pair:    pair,
	}
}

func (msg *MsgSetPair) Route() string {
	return RouterKey
}

func (msg *MsgSetPair) Type() string {
	return TypeMsgSetPair
}

func (msg *MsgSetPair) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetPair) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPair) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
