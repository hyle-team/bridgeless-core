package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateToken = "update_token"

var _ sdk.Msg = &MsgUpdateToken{}

func NewMsgUpdateToken(creator string, tokenId uint64, metadata TokenMetadata) *MsgUpdateToken {
	return &MsgUpdateToken{
		Creator:  creator,
		TokenId:  tokenId,
		Metadata: metadata,
	}
}

func (msg *MsgUpdateToken) Route() string {
	return RouterKey
}

func (msg *MsgUpdateToken) Type() string {
	return TypeMsgUpdateToken
}

func (msg *MsgUpdateToken) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(errorsmod.Wrap(err, "failed to get signers"))
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgUpdateToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return ferrorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", err)
	}

	if err = validateTokenMetadata(&msg.Metadata); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}
