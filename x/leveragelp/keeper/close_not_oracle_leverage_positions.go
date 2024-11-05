package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/amm/types"
)

/*
SPIKE DEV-2106:
We need to close all leverage LP positions from pools that are not oracle and remove
those pools as leverage LP and perpetual
*/
func (k Keeper) CloseNotOracleLeveragePositions(ctx sdk.Context) {
	ammPools := k.amm.GetAllPool(ctx)

	listAmmPools := map[uint64]types.Pool{}

	for _, ammPool := range ammPools {
		listAmmPools[ammPool.PoolId] = ammPool
	}

	leveragePositions := k.GetAllPositions(ctx)

	for _, position := range leveragePositions {
		poolId := position.AmmPoolId
		if _, ok := listAmmPools[poolId]; ok {
			pool := listAmmPools[poolId]
			if !pool.PoolParams.UseOracle {
				/*
					How should we close these positions?
					Should we use "func (k msgServer) ClosePositions" to close the positions?
				*/
			}
		}
	}
}
