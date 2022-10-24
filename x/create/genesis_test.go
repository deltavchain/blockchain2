package create_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hacker/testutil/keeper"
	"hacker/testutil/nullify"
	"hacker/x/create"
	"hacker/x/create/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CreateKeeper(t)
	create.InitGenesis(ctx, *k, genesisState)
	got := create.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
