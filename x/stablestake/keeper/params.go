package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/stablestake/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)

	b := store.Get(types.ParamKeyPrefix)
	if b == nil {
		return
	}

	k.cdc.MustUnmarshal(b, &params)
	return
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&params)
	store.Set(types.ParamKeyPrefix, b)
}

func (k Keeper) GetDepositDenom(ctx sdk.Context) string {
	params := k.GetParams(ctx)
	entry, found := k.assetProfileKeeper.GetEntry(ctx, params.DepositDenom)
	if !found {
		return params.DepositDenom
	}
	return entry.Denom
}

func (k Keeper) GetRedemptionRate(ctx sdk.Context) sdk.Dec {
	params := k.GetParams(ctx)
	totalShares := k.bk.GetSupply(ctx, types.GetShareDenom())

	if totalShares.Amount.IsZero() {
		return sdk.ZeroDec()
	}

	return params.TotalValue.ToLegacyDec().Quo(totalShares.Amount.ToLegacyDec())
}
