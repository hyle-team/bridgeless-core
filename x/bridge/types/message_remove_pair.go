package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRemovePair = "remove-pair"
)

var _ sdk.Msg = &MsgRemovePair{}

// TODO rename TokenInfo to PairInfo
func NewMsgRemovePair(creator string, sourceChain, destinationChain, address string) *MsgRemovePair {
	return &MsgRemovePair{
		Creator:  creator,
		SrcChain: sourceChain,
		DstChain: destinationChain,
		Address:  address,
	}
}

func (msg *MsgRemovePair) Route() string {
	return RouterKey
}

func (msg *MsgRemovePair) Type() string {
	return TypeMsgRemovePair
}

func (msg *MsgRemovePair) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemovePair) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemovePair) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
