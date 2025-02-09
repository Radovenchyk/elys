package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/amm/types"
	assetprofiletypes "github.com/elys-network/elys/x/assetprofile/types"
	oracletypes "github.com/elys-network/elys/x/oracle/types"
)

func (k Keeper) GetExternalLiquidityRatio(ctx sdk.Context, pool types.Pool, amountDepthInfo []types.AssetAmountDepth) ([]types.PoolAsset, error) {
	updatedAssets := make([]types.PoolAsset, len(pool.PoolAssets))
	copy(updatedAssets, pool.PoolAssets)

	for i, asset := range updatedAssets {
		for _, el := range amountDepthInfo {
			entry, found := k.assetProfileKeeper.GetEntryByDenom(ctx, asset.Token.Denom)
			if !found {
				return nil, assetprofiletypes.ErrAssetProfileNotFound
			}
			if entry.DisplayName == el.Asset {
				price, found := k.oracleKeeper.GetAssetPrice(ctx, el.Asset)
				if !found {
					return nil, fmt.Errorf("asset price not set: %s", el.Asset)
				} else {
					O_Tvl := price.Price.Mul(el.Amount)
					P_Tvl := asset.Token.Amount.ToLegacyDec().Mul(price.Price)

					// Ensure tvl is not zero to avoid division by zero
					if P_Tvl.IsZero() {
						return nil, types.ErrAmountTooLow
					}

					liquidityRatio := LiquidityRatioFromPriceDepth(el.Depth)
					// Ensure tvl is not zero to avoid division by zero
					if liquidityRatio.IsZero() {
						return nil, types.ErrAmountTooLow
					}
					asset.ExternalLiquidityRatio = (O_Tvl.Quo(P_Tvl)).Quo(liquidityRatio)

					if asset.ExternalLiquidityRatio.LT(sdk.OneDec()) {
						asset.ExternalLiquidityRatio = sdk.OneDec()
					}
				}
			}
		}
		updatedAssets[i] = asset
	}
	return updatedAssets, nil
}

func LiquidityRatioFromPriceDepth(depth sdk.Dec) sdk.Dec {
	if depth == sdk.OneDec() {
		return sdk.OneDec()
	}
	sqrt, err := sdk.OneDec().Sub(depth).ApproxSqrt()
	if err != nil {
		panic(err)
	}
	return sdk.OneDec().Sub(sqrt)
}

func (k msgServer) FeedMultipleExternalLiquidity(goCtx context.Context, msg *types.MsgFeedMultipleExternalLiquidity) (*types.MsgFeedMultipleExternalLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	feeder, found := k.oracleKeeper.GetPriceFeeder(ctx, sdk.MustAccAddressFromBech32(msg.Sender))
	if !found {
		return nil, oracletypes.ErrNotAPriceFeeder
	}

	if !feeder.IsActive {
		return nil, oracletypes.ErrPriceFeederNotActive
	}

	for _, el := range msg.Liquidity {
		pool, found := k.GetPool(ctx, el.PoolId)
		if !found {
			return nil, types.ErrInvalidPoolId
		}

		// Get external liquidity ratio for each of the asset separately
		poolAssets, err := k.GetExternalLiquidityRatio(ctx, pool, el.AmountDepthInfo)
		if err != nil {
			return nil, err
		}

		pool.PoolAssets = poolAssets
		k.SetPool(ctx, pool)
	}

	return &types.MsgFeedMultipleExternalLiquidityResponse{}, nil
}
