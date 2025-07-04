package keeper

import (
	"context"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *types.GenesisState {
	return &types.GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params: *DefaultParams(),
	}
}

func (k Keeper) ExportState(ctx context.Context, genState *types.GenesisState) error {
	lockTxStore, err := k.LockTransactionStore.Iterate(ctx, nil)
	if err != nil {
		return err
	}

	genState.LockTransactions, err = lockTxStore.Values()
	if err != nil {
		return err
	}

	pendingMintTxStore, err := k.PendingMintTransactionsMap.Iterate(ctx, nil)
	if err != nil {
		return err
	}

	genState.PendingMintTransactions, err = pendingMintTxStore.Values()
	if err != nil {
		return err
	}

	genState.FirstPendingEthMintTransaction, err = k.GetFirstPendingEthMintTransaction(ctx)
	if err != nil {
		return err
	}

	genState.FirstPendingSolMintTransaction, err = k.GetFirstPendingSolMintTransaction(ctx)
	if err != nil {
		return err
	}

	genState.PendingMintTransactionCount, err = k.PendingMintTransactionCount.Get(ctx)
	if err != nil {
		return err
	}

	burnEventStore, err := k.BurnEvents.Iterate(ctx, nil)
	if err != nil {
		return err
	}

	genState.BurnEvents, err = burnEventStore.Values()
	if err != nil {
		return err
	}

	genState.FirstPendingBurnEvent, err = k.FirstPendingBurnEvent.Get(ctx)
	if err != nil {
		return err
	}

	genState.BurnEventCount, err = k.BurnEventCount.Get(ctx)
	if err != nil {
		return err
	}

	redemptionStore, err := k.Redemptions.Iterate(ctx, nil)
	if err != nil {
		return err
	}

	genState.Redemptions, err = redemptionStore.Values()
	if err != nil {
		return err
	}

	genState.FirstPendingRedemption, err = k.FirstPendingRedemption.Get(ctx)
	if err != nil {
		return err
	}

	genState.FirstRedemptionAwaitingSign, err = k.FirstRedemptionAwaitingSign.Get(ctx)
	if err != nil {
		return err
	}

	genState.Supply, err = k.Supply.Get(ctx)
	if err != nil {
		return err
	}

	genState.FirstPendingStakeTransaction, err = k.FirstPendingStakeTransaction.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
