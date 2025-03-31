package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetNewbiesList = "submit_newbies_list"

var _ sdk.Msg = &MsgSetPartiesList{}

func NewMsgSetNewbiesList(creator string, partiesList []*Party) *MsgSetPartiesList {
	return &MsgSetPartiesList{
		Creator: creator,
		Parties: partiesList,
	}
}

func (msg *MsgSetNewbiesList) Route() string {
	return RouterKey
}

func (msg *MsgSetNewbiesList) Type() string {
	return TypeMsgSetNewbiesList
}

func (msg *MsgSetNewbiesList) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetNewbiesList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetNewbiesList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Parties) == 0 {
		return nil
	}

	return validateNewbiePartiesList(msg.Parties)
}
