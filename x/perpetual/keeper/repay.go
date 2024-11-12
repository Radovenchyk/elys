package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ammtypes "github.com/elys-network/elys/x/amm/types"
	"github.com/elys-network/elys/x/perpetual/types"
)

// Repay ammPool has to be pointer because RemoveFromPoolBalance updates pool assets
func (k Keeper) Repay(ctx sdk.Context, mtp *types.MTP, pool *types.Pool, ammPool *ammtypes.Pool, returnAmount math.Int, payingLiabilities math.Int, closingRatio math.LegacyDec, baseCurrency string) error {
	if returnAmount.IsPositive() {
		returnCoins := sdk.NewCoins(sdk.NewCoin(mtp.CustodyAsset, returnAmount))
		err := k.SendFromAmmPool(ctx, ammPool, mtp.GetAccountAddress(), returnCoins)
		if err != nil {
			return err
		}
	}

	mtp.Liabilities = mtp.Liabilities.Sub(payingLiabilities)

	closingCustodyAmount := mtp.Custody.ToLegacyDec().Mul(closingRatio).TruncateInt()
	mtp.Custody = mtp.Custody.Sub(closingCustodyAmount)

	reducingCollateralAmt := closingRatio.Mul(mtp.Collateral.ToLegacyDec()).TruncateInt()
	mtp.Collateral = mtp.Collateral.Sub(reducingCollateralAmt)

	oldTakeProfitCustody := mtp.TakeProfitCustody
	mtp.TakeProfitCustody = mtp.TakeProfitCustody.Sub(mtp.TakeProfitCustody.ToLegacyDec().Mul(closingRatio).TruncateInt())

	oldTakeProfitLiabilities := mtp.TakeProfitLiabilities
	mtp.TakeProfitLiabilities = mtp.TakeProfitLiabilities.Sub(mtp.TakeProfitLiabilities.ToLegacyDec().Mul(closingRatio).TruncateInt())

	err := pool.UpdateCustody(mtp.CustodyAsset, closingCustodyAmount, false, mtp.Position)
	if err != nil {
		return err
	}

	err = pool.UpdateLiabilities(mtp.LiabilitiesAsset, payingLiabilities, false, mtp.Position)
	if err != nil {
		return err
	}

	err = pool.UpdateCollateral(mtp.CollateralAsset, reducingCollateralAmt, false, mtp.Position)
	if err != nil {
		return err
	}

	err = pool.UpdateTakeProfitLiabilities(mtp.LiabilitiesAsset, oldTakeProfitLiabilities.Sub(mtp.TakeProfitLiabilities), false, mtp.Position)
	if err != nil {
		return err
	}

	err = pool.UpdateTakeProfitCustody(mtp.CustodyAsset, oldTakeProfitCustody.Sub(mtp.TakeProfitCustody), false, mtp.Position)
	if err != nil {
		return err
	}

	// This is for accounting purposes, mtp.Custody gets reduced by borrowInterestPaymentCustody and funding fee. so msg.Amount is greater than mtp.Custody here. So if it's negative it should be closed
	if mtp.Custody.IsZero() || mtp.Custody.IsNegative() {
		err = k.DestroyMTP(ctx, mtp.GetAccountAddress(), mtp.Id)
		if err != nil {
			return err
		}
	} else {
		// update mtp health
		mtp.MtpHealth, err = k.GetMTPHealth(ctx, *mtp, *ammPool, baseCurrency)
		if err != nil {
			return err
		}
		err = k.SetMTP(ctx, mtp)
		if err != nil {
			return err
		}
	}

	k.SetPool(ctx, *pool)

	return nil
}
