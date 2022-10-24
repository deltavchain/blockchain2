package hack_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "hacker/testutil/keeper"
	"hacker/testutil/nullify"
	"hacker/x/hack"
	"hacker/x/hack/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HackKeeper(t)
	hack.InitGenesis(ctx, *k, genesisState)
	got := hack.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
