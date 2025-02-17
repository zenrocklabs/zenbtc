package v4_test

import (
	"testing"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/store/types"
	validation "github.com/Zenrock-Foundation/zrchain/v5/x/validation/module"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/test-go/testify/require"
	v4 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v4"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func TestMigrate(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(validation.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	kvStoreService := runtime.NewKVStoreService(storeKey)
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
	err := pendingTxs.Set(ctx, ptxs)
	require.NoError(t, err)

	createdTxs := []*types.PendingMintTransaction{}
	createPendingMintTx := func(ctx sdk.Context, tx *types.PendingMintTransaction) error {
		createdTxs = append(createdTxs, tx)
		return nil
	}

	require.NoError(t, v4.ChangePendingMintTxChainIdtoCaip2Id(ctx, pendingTxs, createPendingMintTx))

	// Verify that both transactions were created with the correct CAIP-2 chain ID
	require.Len(t, createdTxs, 2)
	for _, tx := range createdTxs {
		require.Equal(t, "eip155:17000", tx.Caip2ChainId)
	}
}

func TestMigrate_Fail(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(validation.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	kvStoreService := runtime.NewKVStoreService(storeKey)
	sb := collections.NewSchemaBuilder(kvStoreService)

	pendingTxs := collections.NewItem(sb, types.PendingMintTransactionsKey, types.PendingMintTransactionsIndex, codec.CollValue[types.PendingMintTransactions](cdc))

	createPendingMintTx := func(ctx sdk.Context, tx *types.PendingMintTransaction) error {
		return nil
	}

	require.NoError(t, v4.ChangePendingMintTxChainIdtoCaip2Id(ctx, pendingTxs, createPendingMintTx))
}
