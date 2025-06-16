package v7

import (
	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

const (
	// TargetChainID is the chain ID of the burn events to be removed.
	TargetChainID = "solana:5eykt4UsFv8P8NJdTREpY1vzqKqZKvdp" // Solana mainnet
	// NewFirstPendingBurnEvent is the new value for the first pending burn event.
	NewFirstPendingBurnEvent = uint64(5)
)

// PurgeInvalidBurnEvents removes all burn events with a specific chain ID and sets the first pending burn event.
func PurgeInvalidBurnEvents(
	ctx sdk.Context,
	burnEvents collections.Map[uint64, types.BurnEvent],
	firstPendingBurnEvent collections.Item[uint64],
) error {
	var keysToRemove []uint64

	// Find all burn events with the target chain ID
	if err := burnEvents.Walk(ctx, nil, func(key uint64, value types.BurnEvent) (stop bool, err error) {
		if value.ChainID == TargetChainID {
			keysToRemove = append(keysToRemove, key)
		}
		return false, nil
	}); err != nil {
		return err
	}

	// Remove the identified burn events
	for _, key := range keysToRemove {
		if err := burnEvents.Remove(ctx, key); err != nil {
			return err
		}
	}

	// Set the new first pending burn event
	return firstPendingBurnEvent.Set(ctx, NewFirstPendingBurnEvent)
}
