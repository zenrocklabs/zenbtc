package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v2"
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

// Migrate1to2 migrates x/zenbtc state from consensus version 1 to 2.
func (m Migrator) Migrate1to2(ctx sdk.Context) error {

	if err := v2.UpdateParams(ctx, m.keeper.Params); err != nil {
		return err
	}

	return nil
}
