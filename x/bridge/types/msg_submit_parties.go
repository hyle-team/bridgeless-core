package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitParties = "submit_parties"

var _ sdk.Msg = &MsgSubmitParties{}

func NewMsgSubmitParties(creator string, partiesList []*Party) *MsgSubmitParties {
	return &MsgSubmitParties{
		Creator: creator,
		Parties: partiesList,
	}
}

func (msg *MsgSubmitParties) Route() string {
	return RouterKey
}

func (msg *MsgSubmitParties) Type() string {
	return TypeMsgSubmitParties
}

func (msg *MsgSubmitParties) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgSubmitParties) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitParties) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Parties) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no parties specified")
	}

	return nil
}
