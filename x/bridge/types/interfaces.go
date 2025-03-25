package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingKeeper interface {
	MaxValidators(ctx sdk.Context) (res uint32)
	GetValidators(ctx sdk.Context, maxRetrieve uint32) (validators []types.Validator)
}
