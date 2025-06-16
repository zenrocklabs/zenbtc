package v7_test

import (
	"testing"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/stretchr/testify/require"

	v7 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v7"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func makeTestEncodingConfig() moduletestutil.TestEncodingConfig {
	return moduletestutil.MakeTestEncodingConfig()
}

func TestPurgeInvalidBurnEvents(t *testing.T) {
	encCfg := makeTestEncodingConfig()
	cdc := encCfg.Codec
	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	kvStoreService := runtime.NewKVStoreService(storeKey)

	schemaBuilder := collections.NewSchemaBuilder(kvStoreService)
	// Use the real codec
	collCdc := codec.CollValue[types.BurnEvent](cdc)

	burnEvents := collections.NewMap(schemaBuilder, collections.NewPrefix(0), "burn_events", collections.Uint64Key, collCdc)
	firstPendingBurnEvent := collections.NewItem(schemaBuilder, collections.NewPrefix(2), "first_pending_burn_event", collections.Uint64Value)

	_, err := schemaBuilder.Build()
	require.NoError(t, err)

	// Sample data
	testData := []types.BurnEvent{
		{Id: 1, ChainID: v7.TargetChainID, DestinationAddr: []byte("addr1"), Amount: 100},
		{Id: 2, ChainID: "another-chain", DestinationAddr: []byte("addr2"), Amount: 200},
		{Id: 3, ChainID: v7.TargetChainID, DestinationAddr: []byte("addr3"), Amount: 300},
		{Id: 4, ChainID: "another-chain-2", DestinationAddr: []byte("addr4"), Amount: 400},
	}

	// Populate burn events store
	for _, item := range testData {
		err := burnEvents.Set(ctx, item.Id, item)
		require.NoError(t, err)
	}

	// Set initial first pending burn event
	err = firstPendingBurnEvent.Set(ctx, uint64(1))
	require.NoError(t, err)

	// Run the migration
	err = v7.PurgeInvalidBurnEvents(ctx, burnEvents, firstPendingBurnEvent)
	require.NoError(t, err)

	// --- Verification ---

	// 1. Verify invalid burn events are removed
	_, err = burnEvents.Get(ctx, 1)
	require.Error(t, err, "Event 1 should have been removed")
	_, err = burnEvents.Get(ctx, 3)
	require.Error(t, err, "Event 3 should have been removed")

	// 2. Verify valid burn events still exist
	event2, err := burnEvents.Get(ctx, 2)
	require.NoError(t, err, "Event 2 should still exist")
	require.Equal(t, testData[1], event2)

	event4, err := burnEvents.Get(ctx, 4)
	require.NoError(t, err, "Event 4 should still exist")
	require.Equal(t, testData[3], event4)

	// 3. Verify the count of remaining items
	iter, err := burnEvents.Iterate(ctx, nil)
	require.NoError(t, err)
	allKeys, err := iter.Keys()
	require.NoError(t, err)
	require.Len(t, allKeys, 2, "There should be 2 burn events remaining")

	// 4. Verify first pending burn event is updated
	newFirstPending, err := firstPendingBurnEvent.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, v7.NewFirstPendingBurnEvent, newFirstPending, "First pending burn event should be updated")
}

func TestPurgeInvalidBurnEvents_EmptyStore(t *testing.T) {
	encCfg := makeTestEncodingConfig()
	cdc := encCfg.Codec
	storeKey := storetypes.NewKVStoreKey(types.ModuleName + "_empty")
	tKey := storetypes.NewTransientStoreKey("transient_test_empty")
	ctx := testutil.DefaultContext(storeKey, tKey)
	kvStoreService := runtime.NewKVStoreService(storeKey)

	schemaBuilder := collections.NewSchemaBuilder(kvStoreService)
	collCdc := codec.CollValue[types.BurnEvent](cdc)

	burnEvents := collections.NewMap(schemaBuilder, collections.NewPrefix(0), "burn_events_empty", collections.Uint64Key, collCdc)
	firstPendingBurnEvent := collections.NewItem(schemaBuilder, collections.NewPrefix(2), "first_pending_burn_event_empty", collections.Uint64Value)

	_, err := schemaBuilder.Build()
	require.NoError(t, err)

	// Set initial first pending burn event (even if burn events are empty)
	err = firstPendingBurnEvent.Set(ctx, uint64(0))
	require.NoError(t, err)

	// Run the migration on an empty store
	err = v7.PurgeInvalidBurnEvents(ctx, burnEvents, firstPendingBurnEvent)
	require.NoError(t, err)

	// Verify burn events store is still empty
	iter, err := burnEvents.Iterate(ctx, nil)
	require.NoError(t, err)
	keys, err := iter.Keys()
	require.NoError(t, err)
	require.Empty(t, keys, "Burn events store should remain empty")

	// Verify first pending burn event is updated
	newFirstPending, err := firstPendingBurnEvent.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, v7.NewFirstPendingBurnEvent, newFirstPending, "First pending burn event should be updated even if the store was empty")
}
