package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetBlacklist = "submit_blacklist"

var _ sdk.Msg = &MsgSetBlacklistParties{}

func NewMsgSetBlacklistParties(creator string, partiesList []*Party) *MsgSetBlacklistParties {
	return &MsgSetBlacklistParties{
		Creator: creator,
		Parties: partiesList,
	}
}

func (msg *MsgSetBlacklistParties) Route() string {
	return RouterKey
}

func (msg *MsgSetBlacklistParties) Type() string {
	return TypeMsgSetParties
}

func (msg *MsgSetBlacklistParties) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetBlacklistParties) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetBlacklistParties) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Parties) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no parties specified")
	}

	return nil
}
