package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInsertChain = "insert_chain"

var _ sdk.Msg = &MsgInsertChain{}

func NewMsgInsertChain(creator string, chain Chain) *MsgInsertChain {
	return &MsgInsertChain{
		Creator: creator,
		Chain:   chain,
	}
}

func (msg *MsgInsertChain) Route() string {
	return RouterKey
}

func (msg *MsgInsertChain) Type() string {
	return TypeMsgInsertChain
}

func (msg *MsgInsertChain) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgInsertChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInsertChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if err = validateChain(&msg.Chain); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}
