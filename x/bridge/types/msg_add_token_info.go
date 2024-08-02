package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddTokenInfo = "add_token_info"

var _ sdk.Msg = &MsgAddTokenInfo{}

func NewMsgAddTokenInfo(creator string, info TokenInfo) *MsgAddTokenInfo {
	return &MsgAddTokenInfo{
		Creator: creator,
		Info:    info,
	}
}

func (msg *MsgAddTokenInfo) Route() string {
	return RouterKey
}

func (msg *MsgAddTokenInfo) Type() string {
	return TypeMsgAddTokenInfo
}

func (msg *MsgAddTokenInfo) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgAddTokenInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddTokenInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if err = validateTokenInfo(&msg.Info, nil); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}
