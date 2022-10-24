package keeper

import (
	"hacker/x/create/types"
)

var _ types.QueryServer = Keeper{}
