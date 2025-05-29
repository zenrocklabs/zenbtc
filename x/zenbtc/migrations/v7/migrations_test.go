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

// makeTestEncodingConfig creates a minimal encoding config for testing.
func makeTestEncodingConfig() moduletestutil.TestEncodingConfig {
	return moduletestutil.MakeTestEncodingConfig()
}

func TestClearBurnEvents(t *testing.T) {
	encCfg := makeTestEncodingConfig()
	cdc := encCfg.Codec
	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test_burn")
	ctx := testutil.DefaultContext(storeKey, tKey)
	kvStoreService := runtime.NewKVStoreService(storeKey)

	schemaBuilder := collections.NewSchemaBuilder(kvStoreService)
	burnEventStoreKey := collections.NewPrefix(0) // Using a prefix for the burn event store

	// Value Codec for types.BurnEvent
	burnEventValueCodec := codec.CollValue[types.BurnEvent](cdc)

	burnEventsStore := collections.NewMap(schemaBuilder, burnEventStoreKey, "burn_events", collections.Uint64Key, burnEventValueCodec)

	_, err := schemaBuilder.Build()
	require.NoError(t, err)

	// Sample data
	testBurnEvents := []struct {
		key  uint64
		data types.BurnEvent
	}{
		{
			key:  1,
			data: types.BurnEvent{Id: 1, TxID: "txhash1", Amount: 100, DestinationAddr: []byte("addr1")},
		},
		{
			key:  2,
			data: types.BurnEvent{Id: 2, TxID: "txhash2", Amount: 200, DestinationAddr: []byte("addr2")},
		},
		{
			key:  3,
			data: types.BurnEvent{Id: 3, TxID: "txhash3", Amount: 300, DestinationAddr: []byte("addr3")},
		},
	}

	// Populate the store
	for _, item := range testBurnEvents {
		err := burnEventsStore.Set(ctx, item.key, item.data)
		require.NoError(t, err)
	}

	// Verify store is populated
	iterBefore, err := burnEventsStore.Iterate(ctx, nil)
	require.NoError(t, err)
	keysBefore, err := iterBefore.Keys()
	require.NoError(t, err)
	require.Len(t, keysBefore, len(testBurnEvents), "Store should be populated before clearing")

	// Run ClearBurnEvents
	err = v7.ClearBurnEvents(ctx, burnEventsStore)
	require.NoError(t, err)

	// Verify store is empty
	iterAfter, err := burnEventsStore.Iterate(ctx, nil)
	require.NoError(t, err)
	keysAfter, err := iterAfter.Keys()
	require.NoError(t, err)
	require.Empty(t, keysAfter, "Store should be empty after ClearBurnEvents")

	// Test with an empty store
	// Create a new context and store for this sub-test to avoid interference
	ctxEmpty := testutil.DefaultContext(storeKey, storetypes.NewTransientStoreKey("transient_test_burn_empty"))
	kvStoreServiceEmpty := runtime.NewKVStoreService(storeKey)
	schemaBuilderEmpty := collections.NewSchemaBuilder(kvStoreServiceEmpty)
	burnEventStoreKeyEmpty := collections.NewPrefix(1) // Different prefix to avoid collision
	burnEventsStoreEmpty := collections.NewMap(schemaBuilderEmpty, burnEventStoreKeyEmpty, "burn_events_empty", collections.Uint64Key, burnEventValueCodec)
	_, err = schemaBuilderEmpty.Build()
	require.NoError(t, err)

	// Run ClearBurnEvents on an empty store
	err = v7.ClearBurnEvents(ctxEmpty, burnEventsStoreEmpty)
	require.NoError(t, err)

	// Verify store is still empty
	iterEmptyAfter, err := burnEventsStoreEmpty.Iterate(ctxEmpty, nil)
	require.NoError(t, err)
	keysEmptyAfter, err := iterEmptyAfter.Keys()
	require.NoError(t, err)
	require.Empty(t, keysEmptyAfter, "Store should remain empty when clearing an already empty store")
}
