package keeper

import (
	"context"
	"fmt"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/perpetual/types"
)

func (k msgServer) UpdateTakeProfitPrice(goCtx context.Context, msg *types.MsgUpdateTakeProfitPrice) (*types.MsgUpdateTakeProfitPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Load existing mtp
	creator := sdk.MustAccAddressFromBech32(msg.Creator)
	mtp, err := k.GetMTP(ctx, creator, msg.Id)
	if err != nil {
		return nil, err
	}

	poolId := mtp.AmmPoolId
	pool, found := k.GetPool(ctx, poolId)
	if !found {
		return nil, errorsmod.Wrap(types.ErrPoolDoesNotExist, fmt.Sprintf("poolId: %d", poolId))
	}

	mtp.TakeProfitPrice = msg.Price
	mtp.TakeProfitCustody = types.CalcMTPTakeProfitCustody(mtp)
	mtp.TakeProfitLiabilities, err = k.CalcMTPTakeProfitLiability(ctx, mtp)
	if err != nil {
		return nil, err
	}

	// All take profit liability has to be in liabilities asset
	err = pool.UpdateTakeProfitLiabilities(mtp.LiabilitiesAsset, mtp.TakeProfitLiabilities, true, mtp.Position)
	if err != nil {
		return nil, err
	}

	// All take profit custody has to be in custody asset
	err = pool.UpdateTakeProfitCustody(mtp.CustodyAsset, mtp.TakeProfitCustody, true, mtp.Position)
	if err != nil {
		return nil, err
	}

	err = k.SetMTP(ctx, &mtp)
	if err != nil {
		return nil, err
	}

	k.SetPool(ctx, pool)

	ammPool, err := k.GetAmmPool(ctx, poolId)
	if err != nil {
		return nil, err
	}

	if k.hooks != nil {
		params := k.GetParams(ctx)
		err = k.hooks.AfterPerpetualPositionModified(ctx, ammPool, pool, creator, params.EnableTakeProfitCustodyLiabilities)
		if err != nil {
			return nil, err
		}
	}

	event := sdk.NewEvent(types.EventOpen,
		sdk.NewAttribute("id", strconv.FormatInt(int64(mtp.Id), 10)),
		sdk.NewAttribute("address", mtp.Address),
		sdk.NewAttribute("take_profit_price", mtp.TakeProfitPrice.String()),
	)
	ctx.EventManager().EmitEvent(event)

	return &types.MsgUpdateTakeProfitPriceResponse{}, nil
}
