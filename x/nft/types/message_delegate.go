package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgDelegate = "delegate"
)

var _ sdk.Msg = &MsgDelegate{}

func NewMsgDelegate(
	id uint32,
	owner string,
	recipient string,
) *MsgDelegate {
	return &MsgDelegate{
		Recipient: recipient,
		Owner:     owner,
		Id:        id,
	}
}

func (msg *MsgDelegate) Route() string {
	return RouterKey
}

func (msg *MsgDelegate) Type() string {
	return TypeMsgDelegate
}

func (msg *MsgDelegate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid Owner address")
	}

	// TODO do call to accumulator to validate amount

	return nil
}
