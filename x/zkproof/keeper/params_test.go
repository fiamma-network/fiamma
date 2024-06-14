package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "fiamma/testutil/keeper"
	"fiamma/x/zkproof/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ZkproofKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
