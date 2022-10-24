package keeper

import (
	"hacker/x/blog/types"
)

var _ types.QueryServer = Keeper{}
