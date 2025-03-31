package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetPartiesList = "submit_parties"

var _ sdk.Msg = &MsgSetPartiesList{}

func NewMsgSetPartiesList(creator string, partiesList []*Party) *MsgSetPartiesList {
	return &MsgSetPartiesList{
		Creator: creator,
		Parties: partiesList,
	}
}

func (msg *MsgSetPartiesList) Route() string {
	return RouterKey
}

func (msg *MsgSetPartiesList) Type() string {
	return TypeMsgSetPartiesList
}

func (msg *MsgSetPartiesList) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetPartiesList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetPartiesList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return validateModuleParties(msg.Parties)
}
