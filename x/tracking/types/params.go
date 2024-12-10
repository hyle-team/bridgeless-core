package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair([]byte(BorrowedEventNameKey), &p.BorrowEventName, validateBorrowedEventName),
		paramtypes.NewParamSetPair([]byte(ContractAddress), &p.ContractAddress, validateContractAddress),
		paramtypes.NewParamSetPair([]byte(LiquidatorAddress), &p.LiquidatorAddress, validateLiquidatorAddress),
		paramtypes.NewParamSetPair([]byte(SenderAddress), &p.SenderAddress, validateSenderAddress),
		paramtypes.NewParamSetPair([]byte(LiquidationEventName), &p.LiquidationEventName, validateLiquidationEventName),
		paramtypes.NewParamSetPair([]byte(OracleAddress), &p.OracleAddress, validateOracleAddress),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}

func validateBorrowedEventName(i interface{}) error {
	name, ok := i.(string)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	if name == "" {
		return sdkerrors.Wrap(fmt.Errorf("empty string"), "borrow event name cannot be an empty string")
	}
	return nil
}

func validateContractAddress(i interface{}) error {
	adr, ok := i.(string)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid type: %T", i)
	}

	if _, err := sdk.AccAddressFromHexUnsafe(adr); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address: %s", err.Error())
	}
	return nil
}

func validateLiquidatorAddress(i interface{}) error {
	liq, ok := i.(string)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	if _, err := sdk.AccAddressFromBech32(liq); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid liquidator address: %s", err.Error())
	}

	return nil
}

func validateSenderAddress(i interface{}) error {
	sender, ok := i.(string)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	if _, err := sdk.AccAddressFromHexUnsafe(sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %s", err.Error())
	}
	return nil
}

func validateLiquidationEventName(i interface{}) error {
	name, ok := i.(string)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	if name == "" {
		return sdkerrors.Wrap(fmt.Errorf("empty string"), "liquidation event name cannot be an empty string")
	}
	return nil
}

func validateOracleAddress(i interface{}) error {
	oracle, ok := i.(string)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	if _, err := sdk.AccAddressFromHexUnsafe(oracle); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid oracle address: %s", err.Error())
	}
	return nil
}
