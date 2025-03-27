package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetGoodbye = "submit_goodbye"

var _ sdk.Msg = &MsgSetGoodbye{}

func NewMsgSetGoodbye(creator string, partiesList []*Party) *MsgSetGoodbye {
	return &MsgSetGoodbye{
		Creator: creator,
		Parties: partiesList,
	}
}

func (msg *MsgSetGoodbye) Route() string {
	return RouterKey
}

func (msg *MsgSetGoodbye) Type() string {
	return TypeMsgSetParties
}

func (msg *MsgSetGoodbye) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetGoodbye) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetGoodbye) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return validateGoodbyePartiesList(msg.Parties)
}
