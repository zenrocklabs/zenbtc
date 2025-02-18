package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v2"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
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

	if err := v2.RemoveBadTestnetState(ctx, m.keeper.PendingMintTransactions, m.keeper.cdc); err != nil {
		return err
	}

	shim := func(ctx sdk.Context, tx *types.PendingMintTransaction) error {
		_, err := m.keeper.CreatePendingMintTransaction(ctx, tx)
		return err
	}

	if err := v2.ChangePendingMintTxChainIdtoCaip2Id(ctx, m.keeper.PendingMintTransactions, shim); err != nil {
		return err
	}

	return nil
}
