package keeper

import (
	"hacker/x/hack/types"
)

var _ types.QueryServer = Keeper{}
