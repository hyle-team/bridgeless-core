package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetGoodbyePartiesList = "submit_goodbye_list"

var _ sdk.Msg = &MsgSetGoodbyePartiesList{}

func NewMsgSetGoodbyeList(creator string, partiesList []*Party) *MsgSetGoodbyePartiesList {
	return &MsgSetGoodbyePartiesList{
		Creator: creator,
		Parties: partiesList,
	}
}

func (msg *MsgSetGoodbyePartiesList) Route() string {
	return RouterKey
}

func (msg *MsgSetGoodbyePartiesList) Type() string {
	return TypeMsgSetGoodbyePartiesList
}

func (msg *MsgSetGoodbyePartiesList) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetGoodbyePartiesList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetGoodbyePartiesList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Parties) == 0 {
		return nil
	}

	return validateGoodbyePartiesList(msg.Parties)
}
