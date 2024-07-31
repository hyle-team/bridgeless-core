package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSetChain = "set-chain"
)

var _ sdk.Msg = &MsgSetChain{}

func NewMsgSetChain(creator string, chain *Chain) *MsgSetChain {
	return &MsgSetChain{
		Creator: creator,
		Chain:   chain,
	}
}

func (msg *MsgSetChain) Route() string {
	return RouterKey
}

func (msg *MsgSetChain) Type() string {
	return TypeMsgSetChain
}

func (msg *MsgSetChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Chain == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "chain cant be nil")
	}
	return nil
}
