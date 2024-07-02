package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	TypeMsgMintTokens = "mint"
)

var _ sdk.Msg = &MsgMintTokens{}

func NewMsgMintTokens(
	unlockTimestamp *timestamppb.Timestamp,
	owner string,
	amount uint64,
	uri string,
) *MsgMintTokens {
	return &MsgMintTokens{
		UnlockTimestemp: unlockTimestamp,
		Owner:           owner,
		Amount:          amount,
		Uri:             uri,
	}
}

func (msg *MsgMintTokens) Route() string {
	return RouterKey
}

func (msg *MsgMintTokens) Type() string {
	return TypeMsgMintTokens
}

func (msg *MsgMintTokens) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintTokens) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errors.Wrap(sdkerrors.ErrInvalidAddress, "invalid Owner address")
	}

	// TODO do call to accumulator to validate amount

	return nil
}
