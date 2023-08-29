package keeper_test

import (
	"testing"

	testkeeper "github.com/enledger/energychain/testutil/keeper"
	"github.com/enledger/energychain/x/energychain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EnergychainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
