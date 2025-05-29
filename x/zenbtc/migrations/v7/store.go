package v7

import (
	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func ClearBurnEvents(ctx sdk.Context, store collections.Map[uint64, types.BurnEvent]) error {
	return store.Walk(ctx, nil, func(key uint64, value types.BurnEvent) (stop bool, err error) {
		if err = store.Remove(ctx, key); err != nil {
			return true, err
		}
		return false, nil
	})
}
