package v3_test

import (
	"testing"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/store/types"
	validation "github.com/Zenrock-Foundation/zrchain/v5/x/validation/module"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/test-go/testify/require"
	v3 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v3"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func TestMigrate(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(validation.AppModuleBasic{})
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

	ptxs.Txs = append(ptxs.Txs, tx)
	ptxs.Txs = append(ptxs.Txs, tx2)
	pendingTxs.Set(ctx, ptxs)

	require.NoError(t, v3.ChangePendingMintTxChainIdtoCaip2Id(ctx, pendingTxs, cdc))

	var res types.PendingMintTransactions
	bz, err := store.Get(types.PendingMintTransactionsKey)
	require.NoError(t, err)
	require.NoError(t, cdc.Unmarshal(bz, &res))

	resTxs, _ := pendingTxs.Get(ctx)

	require.Equal(t, resTxs, res)
}

func TestMigrate_Fail(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(validation.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	kvStoreService := runtime.NewKVStoreService(storeKey)
	store := kvStoreService.OpenKVStore(ctx)
	sb := collections.NewSchemaBuilder(kvStoreService)

	pendingTxs := collections.NewItem(sb, types.PendingMintTransactionsKey, types.PendingMintTransactionsIndex, codec.CollValue[types.PendingMintTransactions](cdc))

	require.NoError(t, v3.ChangePendingMintTxChainIdtoCaip2Id(ctx, pendingTxs, cdc))

	var res types.PendingMintTransactions
	bz, err := store.Get(types.PendingMintTransactionsKey)
	require.NoError(t, err)
	require.NoError(t, cdc.Unmarshal(bz, &res))

	resTxs, _ := pendingTxs.Get(ctx)

	require.Equal(t, resTxs, res)
}
