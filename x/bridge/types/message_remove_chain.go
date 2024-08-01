package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgDeleteChain = "remove-chain"
)

var _ sdk.Msg = &MsgDeleteChain{}

func NewMsgDeleteChain(creator, chainId string,
) *MsgDeleteChain {
	return &MsgDeleteChain{
		Creator: creator,
		ChainId: chainId,
	}
}

func (msg *MsgDeleteChain) Route() string {
	return RouterKey
}

func (msg *MsgDeleteChain) Type() string {
	return TypeMsgDeleteChain
}

func (msg *MsgDeleteChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.ChainId == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "missing chain id")
	}

	return nil
}
