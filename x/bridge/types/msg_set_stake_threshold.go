package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetStakeThreshold = "set_stake_threshold"

var _ sdk.Msg = &MsgSetStakeThreshold{}

func NewMsgSetStakeThreshold(creator string, amount string) *MsgSetStakeThreshold {
	return &MsgSetStakeThreshold{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgSetStakeThreshold) Route() string {
	return RouterKey
}

func (msg *MsgSetStakeThreshold) Type() string {
	return TypeMsgSetParties
}

func (msg *MsgSetStakeThreshold) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (msg *MsgSetStakeThreshold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetStakeThreshold) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	amount, ok := sdk.NewIntFromString(msg.Amount)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid amount (%s)", err)
	}

	if !amount.GT(math.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "threshold amount must be greater than zero")
	}

	return nil
}
