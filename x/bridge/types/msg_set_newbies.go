package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetNewbies = "submit_newbies"

var _ sdk.Msg = &MsgSetNewbies{}

func NewMsgSetNewbies(creator string, partiesList []*Party) *MsgSetNewbies {
	return &MsgSetNewbies{
		Creator: creator,
		Parties: partiesList,
	}
}

func (msg *MsgSetNewbies) Route() string {
	return RouterKey
}

func (msg *MsgSetNewbies) Type() string {
	return TypeMsgSetParties
}

func (msg *MsgSetNewbies) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetNewbies) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetNewbies) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Parties) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no parties specified")
	}

	return nil
}
