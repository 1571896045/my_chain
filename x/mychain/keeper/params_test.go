package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "github.com/1571896045/my_chain/testutil/keeper"
    "github.com/1571896045/my_chain/x/mychain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.MychainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
