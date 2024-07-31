package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRemoveChain = "remove-chain"
)

var _ sdk.Msg = &MsgRemoveChain{}

func NewMsgRemoveChain(creator, chainId string,
) *MsgRemoveChain {
	return &MsgRemoveChain{
		Creator: creator,
		ChainId: chainId,
	}
}

func (msg *MsgRemoveChain) Route() string {
	return RouterKey
}

func (msg *MsgRemoveChain) Type() string {
	return TypeMsgRemoveChain
}

func (msg *MsgRemoveChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.ChainId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing chain id")
	}

	return nil
}
