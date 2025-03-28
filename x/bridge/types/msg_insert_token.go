package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInsertToken = "insert_token"

var _ sdk.Msg = &MsgInsertToken{}

func NewMsgInsertToken(creator string, token Token) *MsgInsertToken {
	return &MsgInsertToken{
		Creator: creator,
		Token:   token,
	}
}

func (msg *MsgInsertToken) Route() string {
	return RouterKey
}

func (msg *MsgInsertToken) Type() string {
	return TypeMsgInsertToken
}

func (msg *MsgInsertToken) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgInsertToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInsertToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if err = validateToken(&msg.Token); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}
