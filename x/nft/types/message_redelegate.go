package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRedelegate = "redelegate"
)

var _ sdk.Msg = &MsgRedelegate{}

func NewMsgRedelegate(creator, validatorNew, validatorSrc, address string,
) *MsgRedelegate {
	return &MsgRedelegate{
		Creator:      creator,
		ValidatorNew: validatorNew,
		ValidatorSrc: validatorSrc,
		Address:      address,
	}
}

func (msg *MsgRedelegate) Route() string {
	return RouterKey
}

func (msg *MsgRedelegate) Type() string {
	return TypeMsgRedelegate
}

func (msg *MsgRedelegate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRedelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRedelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.ValidatorSrc)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.ValidatorNew)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid new validator address (%s)", err)
	}

	return nil
}
