package v3_test

import (
	"testing"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/test-go/testify/require"
	v3 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v3"
	zenbtc "github.com/zenrocklabs/zenbtc/x/zenbtc/module"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func TestMigrate(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(zenbtc.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	kvStoreService := runtime.NewKVStoreService(storeKey)
	store := kvStoreService.OpenKVStore(ctx)
	sb := collections.NewSchemaBuilder(kvStoreService)

	pendingTxs := collections.NewItem(sb, types.PendingMintTransactionsKey, types.PendingMintTransactionsIndex, codec.CollValue[types.PendingMintTransactions](cdc))

	ptxs := types.PendingMintTransactions{Txs: []*types.PendingMintTransaction{}}

	tx := &types.PendingMintTransaction{
		ChainId:          17000,
		ChainType:        types.WalletType_WALLET_TYPE_EVM,
		RecipientAddress: "0x123",
		Amount:           0,
		Creator:          "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
		KeyId:            123,
	}

	tx2 := &types.PendingMintTransaction{
		ChainId:          17000,
		ChainType:        types.WalletType_WALLET_TYPE_EVM,
		RecipientAddress: "0x123",
		Amount:           0,
		Creator:          "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
		KeyId:            123,
	}

	tx3 := &types.PendingMintTransaction{ // this tx should be removed
		ChainId:          1,
		ChainType:        types.WalletType_WALLET_TYPE_EVM,
		RecipientAddress: "0x123",
		Amount:           0,
		Creator:          "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
		KeyId:            123,
	}

	ptxs.Txs = append(ptxs.Txs, tx)
	ptxs.Txs = append(ptxs.Txs, tx2)
	ptxs.Txs = append(ptxs.Txs, tx3)
	err := pendingTxs.Set(ctx, ptxs)
	require.NoError(t, err)

	require.NoError(t, v3.RemoveBadTestnetState(ctx, pendingTxs, cdc))

	var res types.PendingMintTransactions
	bz, err := store.Get(types.PendingMintTransactionsKey)
	require.NoError(t, err)
	require.NoError(t, cdc.Unmarshal(bz, &res))

	resTxs, err := pendingTxs.Get(ctx)
	require.NoError(t, err)

	require.Equal(t, len(ptxs.Txs)-1, len(resTxs.Txs))
}

func TestMigrate_Fail(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(zenbtc.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	kvStoreService := runtime.NewKVStoreService(storeKey)
	store := kvStoreService.OpenKVStore(ctx)
	sb := collections.NewSchemaBuilder(kvStoreService)

	pendingTxs := collections.NewItem(sb, types.PendingMintTransactionsKey, types.PendingMintTransactionsIndex, codec.CollValue[types.PendingMintTransactions](cdc))

	require.NoError(t, v3.RemoveBadTestnetState(ctx, pendingTxs, cdc))

	var res types.PendingMintTransactions
	bz, err := store.Get(types.PendingMintTransactionsKey)
	require.NoError(t, err)
	require.NoError(t, cdc.Unmarshal(bz, &res))

	require.Equal(t, types.PendingMintTransactions{}, res)
}
