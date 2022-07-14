package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/neutron-org/neutron/testutil/interchaintxs/keeper"
	"github.com/neutron-org/neutron/x/interchaintxs/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.InterchainTxsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
