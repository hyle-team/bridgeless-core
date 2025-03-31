package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetBlacklistPartiesList = "submit_blacklist_list"

var _ sdk.Msg = &MsgSetBlacklistPartiesList{}

func NewMsgSetBlacklistPartiesList(creator string, partiesList []*Party) *MsgSetBlacklistPartiesList {
	return &MsgSetBlacklistPartiesList{
		Creator: creator,
		Parties: partiesList,
	}
}

func (msg *MsgSetBlacklistPartiesList) Route() string {
	return RouterKey
}

func (msg *MsgSetBlacklistPartiesList) Type() string {
	return TypeMsgSetBlacklistPartiesList
}

func (msg *MsgSetBlacklistPartiesList) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetBlacklistPartiesList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetBlacklistPartiesList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return validateBlacklistPartiesList(msg.Parties)
}
