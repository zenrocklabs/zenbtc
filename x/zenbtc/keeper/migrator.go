package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v2"
	v3 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v3"
	v4 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v4"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper *Keeper
}

// NewMigrator returns a new Migrator instance.
func NewMigrator(keeper *Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}

// Migrate1to2 migrates x/zenbtc params from consensus version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {

	if err := v2.UpdateParams(ctx, m.keeper.Params); err != nil {
		return err
	}

	return nil
}

// Migrate2to3 migrates x/zenbtc params from consensus version 2 to 3.
func (m Migrator) Migrate2to3(ctx sdk.Context) error {

	if err := v3.UpdateParams(ctx, m.keeper.Params); err != nil {
		return err
	}

	return nil
}

// Migrate3to4 migrates x/zenbtc params from consensus version 3 to 4.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {

	if err := v4.UpdateParams(ctx, m.keeper.Params); err != nil {
		return err
	}

	return nil
}
