package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/Zenrock-Foundation/zrchain/v5/testutil/keeper"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ZenbtcKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
