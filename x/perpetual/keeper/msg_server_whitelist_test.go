package keeper_test

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/elys-network/elys/testutil/sample"
	"github.com/elys-network/elys/x/perpetual/keeper"
	"github.com/elys-network/elys/x/perpetual/types"
)

func (suite *PerpetualKeeperTestSuite) TestMsgServerWhileList() {

	t := suite.T()
	t.Parallel()

	t.Run("ErrInvalidSigner", func(t *testing.T) {
		msg := keeper.NewMsgServerImpl(*suite.app.PerpetualKeeper)
		_, err := msg.Whitelist(suite.ctx, &types.MsgWhitelist{
			Authority:          sample.AccAddress(),
			WhitelistedAddress: "",
		})
		suite.Require().ErrorIs(err, govtypes.ErrInvalidSigner)
	})

	t.Run("ErrWhitelistedAddress", func(t *testing.T) {
		msg := keeper.NewMsgServerImpl(*suite.app.PerpetualKeeper)
		_, err := msg.Whitelist(suite.ctx, &types.MsgWhitelist{
			Authority:          authtypes.NewModuleAddress(govtypes.ModuleName).String(),
			WhitelistedAddress: "error",
		})
		suite.Require().Error(err)
		suite.Require().Contains(err.Error(), "decoding bech32 failed:")
	})

	t.Run("Successful", func(t *testing.T) {
		msg := keeper.NewMsgServerImpl(*suite.app.PerpetualKeeper)
		_, err := msg.Whitelist(suite.ctx, &types.MsgWhitelist{
			Authority:          authtypes.NewModuleAddress(govtypes.ModuleName).String(),
			WhitelistedAddress: sample.AccAddress(),
		})
		suite.Require().Nil(err)
	})

}
