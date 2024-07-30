package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitTransaction = "submit_transaction"

var _ sdk.Msg = &MsgSubmitTransaction{}

func NewMsgSubmitTransaction(submitter string, transaction Transaction) *MsgSubmitTransaction {
	return &MsgSubmitTransaction{
		Submitter:   submitter,
		Transaction: transaction,
	}
}

func (msg *MsgSubmitTransaction) Route() string {
	return RouterKey
}

func (msg *MsgSubmitTransaction) Type() string {
	return TypeMsgSubmitTransaction
}

func (msg *MsgSubmitTransaction) GetSigners() []sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddress}
}

func (msg *MsgSubmitTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid submitter address: %s", err)
	}

	if err = validateTransaction(&msg.Transaction); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}
