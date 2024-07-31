package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRemoveToken = "remove-token"
)

var _ sdk.Msg = &MsgRemoveToken{}

func NewMsgRemoveToken(creator string, tokenId uint64) *MsgRemoveToken {
	return &MsgRemoveToken{
		Creator: creator,
		TokenId: tokenId,
	}
}

func (msg *MsgRemoveToken) Route() string {
	return RouterKey
}

func (msg *MsgRemoveToken) Type() string {
	return TypeMsgRemoveToken
}

func (msg *MsgRemoveToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
