package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v2 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v2"
	v3 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v3"
	v4 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v4"
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
	return v2.UpdateParams(ctx, m.keeper.Params)
}

func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.RemoveBadTestnetState(ctx, m.keeper.PendingMintTransactions, m.keeper.cdc)
}

func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	shim := func(ctx sdk.Context, tx *types.PendingMintTransaction) error {
		_, err := m.keeper.CreatePendingMintTransaction(ctx, tx)
		return err
	}
	return v4.ChangePendingMintTxChainIdtoCaip2Id(ctx, m.keeper.PendingMintTransactions, shim)
}
