package mychain_test

import (
	"testing"

	keepertest "github.com/1571896045/my_chain/testutil/keeper"
	"github.com/1571896045/my_chain/testutil/nullify"
	mychain "github.com/1571896045/my_chain/x/mychain/module"
	"github.com/1571896045/my_chain/x/mychain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MychainKeeper(t)
	mychain.InitGenesis(ctx, k, genesisState)
	got := mychain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
