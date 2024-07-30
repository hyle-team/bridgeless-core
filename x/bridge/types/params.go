package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	_ paramtypes.ParamSet = (*Params)(nil)
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyPrefix(ParamEvmAdminKey), &p.EvmAdmin, validateModuleAdmin),
	}
}

// NewParams creates a new Params instance
func NewParams(evmAdmin string) Params {
	return Params{
		EvmAdmin: evmAdmin,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("cosmos1v4xjvmka36peq6lp2vahxax96540nek53cfyrq")
}

// Validate validates the set of params
func (p Params) Validate() error {
	return validateModuleAdmin(p.EvmAdmin)
}

func validateModuleAdmin(i interface{}) error {
	adm, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	_, err := sdk.AccAddressFromBech32(adm)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	return nil
}
