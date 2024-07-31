package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSetToken = "set-token"
)

var _ sdk.Msg = &MsgSetToken{}

func NewMsgSetToken(creator string, token *Token) *MsgSetToken {
	return &MsgSetToken{
		Creator: creator,
		Token:   token,
	}
}

func (msg *MsgSetToken) Route() string {
	return RouterKey
}

func (msg *MsgSetToken) Type() string {
	return TypeMsgSetToken
}

func (msg *MsgSetToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Token == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "token cannot be nil")
	}
	return nil
}
