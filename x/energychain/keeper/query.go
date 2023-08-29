package keeper

import (
	"github.com/enledger/energychain/x/energychain/types"
)

var _ types.QueryServer = Keeper{}
