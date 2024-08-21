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
		paramtypes.NewParamSetPair([]byte(ParamEvmAdminKey), &p.EvmAdmin, validateModuleAdmin),
		paramtypes.NewParamSetPair([]byte(ParamModuleAdminKey), &p.ModuleAdmin, validateModuleAdmin),
	}
}

// NewParams creates a new Params instance
func NewParams(moduleAdmin, evmAdmin string) Params {
	return Params{
		ModuleAdmin: moduleAdmin,
		EvmAdmin:    evmAdmin,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("", "")
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateModuleAdmin(p.EvmAdmin); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid evm admin address (%s)", err)
	}

	if err := validateModuleAdmin(p.ModuleAdmin); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid module admin address (%s)", err)
	}

	return nil
}

func validateModuleAdmin(i interface{}) error {
	adm, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	_, err := sdk.AccAddressFromBech32(adm)
	if err != nil {
		return fmt.Errorf("failed to parse address: %w", err)
	}

	return nil
}
