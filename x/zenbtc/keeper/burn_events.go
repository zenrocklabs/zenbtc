package keeper

import (
	"errors"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func (k Keeper) CreateBurnEvent(ctx sdk.Context, burnEvent *types.BurnEvent) (uint64, error) {
	count, err := k.BurnEventCount.Get(ctx)
	if err != nil {
		if !errors.Is(err, collections.ErrNotFound) {
			return 0, err
		} else {
			burnEvent.Id = 1
		}
	} else {
		burnEvent.Id = count + 1
	}

	burnEvent.Status = types.BurnStatus_BURN_STATUS_BURNED

	if err := k.BurnEvents.Set(ctx, burnEvent.Id, *burnEvent); err != nil {
		return 0, err
	}

	if err := k.BurnEventCount.Set(ctx, burnEvent.Id); err != nil {
		return 0, err
	}

	return burnEvent.Id, nil
}
