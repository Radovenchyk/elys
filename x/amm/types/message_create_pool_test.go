package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/elys-network/elys/testutil/sample"
	"github.com/elys-network/elys/x/amm/types"
	ptypes "github.com/elys-network/elys/x/parameter/types"
	"github.com/stretchr/testify/require"
)

func TestMsgCreatePool_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreatePool
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgCreatePool{
				Sender: "invalid_address",
				PoolParams: &types.PoolParams{
					SwapFee:                     sdk.ZeroDec(),
					ExitFee:                     sdk.ZeroDec(),
					UseOracle:                   false,
					WeightBreakingFeeMultiplier: sdk.ZeroDec(),
					WeightBreakingFeeExponent:   sdk.NewDecWithPrec(25, 1), // 2.5
					WeightRecoveryFeePortion:    sdk.NewDecWithPrec(10, 2), // 10%
					ThresholdWeightDifference:   sdk.ZeroDec(),
					FeeDenom:                    ptypes.BaseCurrency,
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "pool assets must be exactly two",
			msg: types.MsgCreatePool{
				Sender: sample.AccAddress(),
				PoolParams: &types.PoolParams{
					SwapFee:                     sdk.ZeroDec(),
					ExitFee:                     sdk.ZeroDec(),
					UseOracle:                   false,
					WeightBreakingFeeMultiplier: sdk.ZeroDec(),
					WeightBreakingFeeExponent:   sdk.NewDecWithPrec(25, 1), // 2.5
					WeightRecoveryFeePortion:    sdk.NewDecWithPrec(10, 2), // 10%
					ThresholdWeightDifference:   sdk.ZeroDec(),
					FeeDenom:                    ptypes.BaseCurrency,
				},
				PoolAssets: []types.PoolAsset{
					{
						Token:  sdk.NewCoin("uatom", sdk.NewInt(10000000)),
						Weight: sdk.NewInt(10),
					},
				},
			},
			err: types.ErrPoolAssetsMustBeTwo,
		},
		{
			name: "valid address",
			msg: types.MsgCreatePool{
				Sender: sample.AccAddress(),
				PoolParams: &types.PoolParams{
					SwapFee:                     sdk.ZeroDec(),
					ExitFee:                     sdk.ZeroDec(),
					UseOracle:                   false,
					WeightBreakingFeeMultiplier: sdk.ZeroDec(),
					WeightBreakingFeeExponent:   sdk.NewDecWithPrec(25, 1), // 2.5
					WeightRecoveryFeePortion:    sdk.NewDecWithPrec(10, 2), // 10%
					ThresholdWeightDifference:   sdk.ZeroDec(),
					FeeDenom:                    ptypes.BaseCurrency,
				},
				PoolAssets: []types.PoolAsset{
					{
						Token:  sdk.NewCoin("uusdc", sdk.NewInt(10000000)),
						Weight: sdk.NewInt(10),
					},
					{
						Token:  sdk.NewCoin("uatom", sdk.NewInt(10000000)),
						Weight: sdk.NewInt(10),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
