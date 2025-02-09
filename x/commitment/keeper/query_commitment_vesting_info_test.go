package keeper_test

import (
	"testing"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/elys-network/elys/app"
	simapp "github.com/elys-network/elys/app"

	"github.com/elys-network/elys/x/commitment/types"
	ptypes "github.com/elys-network/elys/x/parameter/types"
)

// TestKeeper_CommitmentVestingInfo tests the CommitmentVestingInfo method
func TestKeeper_CommitmentVestingInfo(t *testing.T) {
	app := app.InitElysTestApp(true)
	ctx := app.BaseApp.NewContext(true, tmproto.Header{})

	mk := app.CommitmentKeeper

	// Generate 1 random account with 1000stake balanced
	addr := simapp.AddTestAddrs(app, ctx, 1, sdk.NewInt(1000000000000))

	// Define the test data
	creator := addr[0]

	// Set up the commitments for the creator
	commitments := types.Commitments{
		Creator: creator.String(),
		VestingTokens: []*types.VestingTokens{
			{
				Denom:                ptypes.Eden,
				TotalAmount:          sdk.NewInt(100),
				ClaimedAmount:        sdk.NewInt(50),
				NumBlocks:            10,
				StartBlock:           1,
				VestStartedTimestamp: 1,
			},
		},
	}

	mk.SetCommitments(ctx, commitments)

	actualRes, err := mk.CommitmentVestingInfo(ctx, &types.QueryCommitmentVestingInfoRequest{
		Address: creator.String(),
	})
	require.NoError(t, err)
	require.NotNil(t, actualRes)

	expectedRes := &types.QueryCommitmentVestingInfoResponse{
		Total: sdk.NewInt(50),
		VestingDetails: []types.VestingDetails{
			{
				Id:              "0",
				TotalVesting:    sdk.NewInt(100),
				Claimed:         sdk.NewInt(50),
				VestedSoFar:     sdk.NewInt(-10),
				RemainingBlocks: 11,
			},
		},
	}

	require.Equal(t, expectedRes, actualRes)
}

// TestKeeper_CommitmentVestingInfoNilRequest tests the CommitmentVestingInfo method with nil request
func TestKeeper_CommitmentVestingInfoNilRequest(t *testing.T) {
	app := app.InitElysTestApp(true)
	ctx := app.BaseApp.NewContext(true, tmproto.Header{})

	mk := app.CommitmentKeeper

	_, err := mk.CommitmentVestingInfo(ctx, nil)
	require.Error(t, err)
}
