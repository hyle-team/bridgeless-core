package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgRemovePairById = "delete-pair-id"
)

var _ sdk.Msg = &MsgRemovePairById{}

// TODO rename TokenInfo to PairInfo
func NewMsgRemovePairById(creator string, chainId string, tokenId uint64) *MsgRemovePairById {
	return &MsgRemovePairById{
		Creator: creator,
		ChainId: chainId,
		TokenId: tokenId,
	}
}

func (msg *MsgRemovePairById) Route() string {
	return RouterKey
}

func (msg *MsgRemovePairById) Type() string {
	return TypeMsgRemovePairById
}

func (msg *MsgRemovePairById) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemovePairById) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemovePairById) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
