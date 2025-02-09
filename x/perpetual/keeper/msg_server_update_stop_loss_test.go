package keeper_test

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/testutil/sample"
	ptypes "github.com/elys-network/elys/x/parameter/types"
	"github.com/elys-network/elys/x/perpetual/keeper"
	"github.com/elys-network/elys/x/perpetual/types"
)

func (suite *PerpetualKeeperTestSuite) TestMsgServerUpdateStopLoss_ErrorGetMtp() {
	k := suite.app.PerpetualKeeper
	ctx := suite.ctx
	msg := keeper.NewMsgServerImpl(*k)
	_, err := msg.UpdateStopLoss(ctx, &types.MsgUpdateStopLoss{
		Creator: sample.AccAddress(),
		Id:      1,
		Price:   math.LegacyMustNewDecFromStr("0.98"),
	})
	suite.Require().ErrorIs(err, types.ErrMTPDoesNotExist)
}

func (suite *PerpetualKeeperTestSuite) TestMsgServerUpdateStopLoss_ErrPoolDoesNotExist() {
	k := suite.app.PerpetualKeeper
	ctx := suite.ctx
	msg := keeper.NewMsgServerImpl(*k)

	firstPool := uint64(1)
	suite.SetPerpetualPool(firstPool)

	amount := math.NewInt(400)
	addr := suite.AddAccounts(2, nil)
	firstPositionCreator := addr[0]

	firstOpenPositionMsg := &types.MsgOpen{
		Creator:         firstPositionCreator.String(),
		Leverage:        math.LegacyNewDec(5),
		Position:        types.Position_SHORT,
		PoolId:          firstPool,
		TradingAsset:    ptypes.ATOM,
		Collateral:      sdk.NewCoin(ptypes.BaseCurrency, amount),
		TakeProfitPrice: math.LegacyMustNewDecFromStr("0.95"),
		StopLossPrice:   sdk.ZeroDec(),
	}

	firstPosition, err := suite.app.PerpetualKeeper.Open(ctx, firstOpenPositionMsg, false)
	suite.Require().NoError(err)
	k.RemovePool(ctx, firstPool)
	_, err = msg.UpdateStopLoss(ctx, &types.MsgUpdateStopLoss{
		Creator: firstPositionCreator.String(),
		Id:      firstPosition.Id,
		Price:   math.LegacyMustNewDecFromStr("0.98"),
	})
	suite.Require().ErrorIs(err, types.ErrPoolDoesNotExist)
}

func (suite *PerpetualKeeperTestSuite) TestMsgServerUpdateStopLoss_Successful() {
	k := suite.app.PerpetualKeeper
	ctx := suite.ctx
	msg := keeper.NewMsgServerImpl(*k)

	firstPool := uint64(1)
	suite.SetPerpetualPool(firstPool)

	amount := math.NewInt(400)
	addr := suite.AddAccounts(2, nil)
	firstPositionCreator := addr[0]

	firstOpenPositionMsg := &types.MsgOpen{
		Creator:         firstPositionCreator.String(),
		Leverage:        math.LegacyNewDec(5),
		Position:        types.Position_SHORT,
		PoolId:          firstPool,
		TradingAsset:    ptypes.ATOM,
		Collateral:      sdk.NewCoin(ptypes.BaseCurrency, amount),
		TakeProfitPrice: math.LegacyMustNewDecFromStr("0.95"),
		StopLossPrice:   sdk.ZeroDec(),
	}

	firstPosition, err := suite.app.PerpetualKeeper.Open(ctx, firstOpenPositionMsg, false)
	suite.Require().NoError(err)
	_, err = msg.UpdateStopLoss(ctx, &types.MsgUpdateStopLoss{
		Creator: firstPositionCreator.String(),
		Id:      firstPosition.Id,
		Price:   math.LegacyMustNewDecFromStr("0.98"),
	})
	suite.Require().Nil(err)
}
