package v2_test

import (
	"testing"

	"cosmossdk.io/collections"
	validation "github.com/Zenrock-Foundation/zrchain/v5/x/validation/module"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"
	v2 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v2"
	zenbtc "github.com/zenrocklabs/zenbtc/x/zenbtc/module"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

var (
	params = &types.Params{
		EthBatcherAddr:      "0x912D79F8d489d0d007aBE0E26fD5d2f06BA4A2AA",
		DepositKeyringAddr:  "keyring1hpyh7xqr2w7h4eas5y8twnsg",
		WithdrawerKeyID:     1,
		EthMinterKeyID:      2,
		ChangeAddressKeyIDs: []uint64{3},
		UnstakerKeyID:       4,
		RewardsDepositKeyID: 5,
		BitcoinProxyAddress: "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
	}
)

func TestMigrate1(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(validation.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	kvStoreService := runtime.NewKVStoreService(storeKey)
	store := kvStoreService.OpenKVStore(ctx)
	sb := collections.NewSchemaBuilder(kvStoreService)
	params := collections.NewItem(sb, types.ParamsKey, types.ParamsIndex, codec.CollValue[types.Params](cdc))
	err := params.Set(ctx, types.Params{})
	require.NoError(t, err)
	require.NoError(t, v2.UpdateParams(ctx, params))

	var res types.Params
	bz, err := store.Get(types.ParamsKey)
	require.NoError(t, err)
	require.NoError(t, cdc.Unmarshal(bz, &res))

	require.Equal(t, params, res)
}

func TestMigrateFail1(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(validation.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	kvStoreService := runtime.NewKVStoreService(storeKey)
	sb := collections.NewSchemaBuilder(kvStoreService)
	params := collections.NewItem(sb, types.ParamsKey, types.ParamsIndex, codec.CollValue[types.Params](cdc))

	require.NoError(t, v2.UpdateParams(ctx, params))
}

// type mockSubspace struct {
// 	ps types.HVParams
// }

// func newMockSubspace(ps types.HVParams) mockSubspace {
// 	return mockSubspace{ps: ps}
// }

// func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps exported.ParamSet) {
// 	*ps.(*types.Params) = ms.ps
// }

func TestMigrate2(t *testing.T) {
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

	require.NoError(t, v2.RemoveBadTestnetState(ctx, pendingTxs, cdc))

	var res types.PendingMintTransactions
	bz, err := store.Get(types.PendingMintTransactionsKey)
	require.NoError(t, err)
	require.NoError(t, cdc.Unmarshal(bz, &res))

	resTxs, err := pendingTxs.Get(ctx)
	require.NoError(t, err)

	require.Equal(t, len(ptxs.Txs)-1, len(resTxs.Txs))
}

func TestMigrate_Fail2(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(zenbtc.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	kvStoreService := runtime.NewKVStoreService(storeKey)
	store := kvStoreService.OpenKVStore(ctx)
	sb := collections.NewSchemaBuilder(kvStoreService)

	pendingTxs := collections.NewItem(sb, types.PendingMintTransactionsKey, types.PendingMintTransactionsIndex, codec.CollValue[types.PendingMintTransactions](cdc))

	require.NoError(t, v2.RemoveBadTestnetState(ctx, pendingTxs, cdc))

	var res types.PendingMintTransactions
	bz, err := store.Get(types.PendingMintTransactionsKey)
	require.NoError(t, err)
	require.NoError(t, cdc.Unmarshal(bz, &res))

	require.Equal(t, types.PendingMintTransactions{}, res)
}

func TestMigrate3(t *testing.T) {
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

	require.NoError(t, v2.ChangePendingMintTxChainIdtoCaip2Id(ctx, pendingTxs, createPendingMintTx))

	// Verify that both transactions were created with the correct CAIP-2 chain ID
	require.Len(t, createdTxs, 2)
	for _, tx := range createdTxs {
		require.Equal(t, "eip155:17000", tx.Caip2ChainId)
	}
}

func TestMigrate_Fail3(t *testing.T) {
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

	require.NoError(t, v2.ChangePendingMintTxChainIdtoCaip2Id(ctx, pendingTxs, createPendingMintTx))
}
