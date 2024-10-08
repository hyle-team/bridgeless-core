package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/hyle-team/bridgeless-core/x/bridge/migrations/v2"
	v3 "github.com/hyle-team/bridgeless-core/x/bridge/migrations/v3"
)

type Migrator struct {
	keeper Keeper
}

func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	if err := v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc); err != nil {
		return err
	}

	return nil
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := v3.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc); err != nil {
		return err
	}

	return nil
}
