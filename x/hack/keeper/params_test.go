package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "hacker/testutil/keeper"
	"hacker/x/hack/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HackKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
