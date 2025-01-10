package v6_test

import (
	"testing"

	"cosmossdk.io/collections"
	validation "github.com/Zenrock-Foundation/zrchain/v5/x/validation/module"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"
	v1 "github.com/zenrocklabs/zenbtc/x/zenbtc/migrations/v1"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

var (
	params = &types.Params{
		EthBatcherAddr:      "0x912D79F8d489d0d007aBE0E26fD5d2f06BA4A2AA",
		DepositKeyringAddr:  "keyring1hpyh7xqr2w7h4eas5y8twnsg",
		WithdrawerKeyID:     1,
		MinterKeyID:         2,
		ChangeAddressKeyIDs: []uint64{3},
		UnstakerKeyID:       4,
		RewardsDepositKeyID: 5,
		BitcoinProxyAddress: "zen13y3tm68gmu9kntcxwvmue82p6akacnpt2v7nty",
	}
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
	params := collections.NewItem(sb, types.ParamsKey, types.ParamsIndex, codec.CollValue[types.Params](cdc))
	err := params.Set(ctx, types.Params{})
	require.NoError(t, err)
	require.NoError(t, v1.UpdateParams(ctx, params))

	var res types.Params
	bz, err := store.Get(types.ParamsKey)
	require.NoError(t, err)
	require.NoError(t, cdc.Unmarshal(bz, &res))

	require.Equal(t, params, res)
}

func TestMigrateFail(t *testing.T) {
	encCfg := moduletestutil.MakeTestEncodingConfig(validation.AppModuleBasic{})
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	kvStoreService := runtime.NewKVStoreService(storeKey)
	sb := collections.NewSchemaBuilder(kvStoreService)
	params := collections.NewItem(sb, types.ParamsKey, types.ParamsIndex, codec.CollValue[types.Params](cdc))

	require.Error(t, v1.UpdateParams(ctx, params))
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
