package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"errors"
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
		paramtypes.NewParamSetPair(ParamModuleAdminKey, &p.ModuleAdmin, validateModuleAdmin),
		paramtypes.NewParamSetPair(ParamModulePartiesKey, &p.Parties, validateModuleParties),
		paramtypes.NewParamSetPair(ParamModuleNewbiesKey, &p.Newbies, validateNewbieParties),
		paramtypes.NewParamSetPair(ParamModuleGoodbyeKey, &p.GoodbyeList, validateGoodbyeParties),
		paramtypes.NewParamSetPair(ParamModuleBlacklistKey, &p.Blacklist, validateBlacklistParties),
		paramtypes.NewParamSetPair(ParamModuleStakingThresholdKey, &p.StakeThreshold, validateStakeThreshold),
		paramtypes.NewParamSetPair(ParamModuleTssThresholdKey, &p.TssThreshold, validateTssThreshold),
	}
}

// NewParams creates a new Params instance
func NewParams(moduleAdmin string, parties []*Party) Params {
	return Params{
		ModuleAdmin: moduleAdmin,
		Parties:     parties,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams("", []*Party{})
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateModuleAdmin(p.ModuleAdmin); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid module admin address (%s)", err)
	}

	if err := validateModuleParties(p.Parties); err != nil {
		return errorsmod.Wrapf(ErrInvalidPartiesList, "invalid parties list (%s)", err)
	}

	if err := validateBlacklistParties(p.Parties); err != nil {
		return errorsmod.Wrapf(ErrInvalidPartiesList, "invalid parties list (%s)", err)
	}

	if err := validateGoodbyeParties(p.Parties); err != nil {
		return errorsmod.Wrapf(ErrInvalidPartiesList, "invalid parties list (%s)", err)
	}

	if err := validateNewbieParties(p.Parties); err != nil {
		return errorsmod.Wrapf(ErrInvalidPartiesList, "invalid parties list (%s)", err)
	}

	if err := validateStakeThreshold(p.StakeThreshold); err != nil {
		return errorsmod.Wrap(err, "invalid stake threshold")
	}

	if err := validateTssThreshold(p.TssThreshold); err != nil {
		return errorsmod.Wrap(err, "invalid tss threshold")
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
func validateModuleParties(i interface{}) error {
	parties, ok := i.([]*Party)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	for _, party := range parties {
		_, err := sdk.AccAddressFromBech32(party.Address)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid party address (%s)", err)
		}
	}

	return nil
}

func validateNewbieParties(i interface{}) error {
	return errorsmod.Wrapf(validateModuleParties(i), "invalid newbies parties")
}

func validateGoodbyeParties(i interface{}) error {
	return errorsmod.Wrapf(validateModuleParties(i), "invalid goodbye parties")
}

func validateBlacklistParties(i interface{}) error {
	return errorsmod.Wrapf(validateModuleParties(i), "invalid blacklist parties")
}

func validateThreshold(i interface{}) error {
	n, ok := i.(string)
	if !ok {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidType, "invalid threshold type: %T", i)
	}

	amount, ok := sdk.NewIntFromString(n)
	if !ok {
		return errors.New("invalid threshold amount")
	}

	if !amount.GT(math.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "threshold amount must be greater than zero")
	}

	return nil
}

func validateStakeThreshold(i interface{}) error {
	return errorsmod.Wrap(validateThreshold(i), "invalid stake threshold")
}

func validateTssThreshold(i interface{}) error {
	return errorsmod.Wrap(validateThreshold(i), "invalid Tss threshold")
}
