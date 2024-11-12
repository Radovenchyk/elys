package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSpotOrder{}

func NewMsgCreateSpotOrder(ownerAddress string, orderType SpotOrderType,
	orderPrice OrderPrice, orderAmount sdk.Coin,
	orderTargetDenom string) *MsgCreateSpotOrder {
	return &MsgCreateSpotOrder{
		OrderType:        orderType,
		OrderPrice:       &orderPrice,
		OrderAmount:      &orderAmount,
		OwnerAddress:     ownerAddress,
		OrderTargetDenom: orderTargetDenom,
	}
}

func (msg *MsgCreateSpotOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Validate order price
	if msg.OrderPrice == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "order price cannot be nil")
	}

	// Validate order price
	if msg.OrderPrice.Rate.IsNegative() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "order price cannot be negative")
	}

	err = sdk.ValidateDenom(msg.OrderPrice.BaseDenom)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid base asset denom (%s)", err)
	}

	err = sdk.ValidateDenom(msg.OrderPrice.QuoteDenom)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid quote asset denom (%s)", err)
	}

	// Validate order amount
	if !msg.OrderAmount.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid order amount")
	}

	err = sdk.ValidateDenom(msg.OrderTargetDenom)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid order target asset denom (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateSpotOrder{}

func NewMsgUpdateSpotOrder(creator string, id uint64, orderPrice *OrderPrice) *MsgUpdateSpotOrder {
	return &MsgUpdateSpotOrder{
		OrderId:      id,
		OwnerAddress: creator,
		OrderPrice:   orderPrice,
	}
}

func (msg *MsgUpdateSpotOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	// Validate order price
	if msg.OrderPrice == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "order price cannot be nil")
	}

	// Validate order price
	if msg.OrderPrice.Rate.IsNegative() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "order price cannot be negative")
	}

	err = sdk.ValidateDenom(msg.OrderPrice.BaseDenom)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid base asset denom (%s)", err)
	}

	err = sdk.ValidateDenom(msg.OrderPrice.QuoteDenom)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid quote asset denom (%s)", err)
	}

	// Validate order ID
	if msg.OrderId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "order price cannot be 0")
	}
	return nil
}

var _ sdk.Msg = &MsgCancelSpotOrder{}

func NewMsgCancelSpotOrder(ownerAddress string, orderId uint64) *MsgCancelSpotOrder {
	return &MsgCancelSpotOrder{
		OwnerAddress: ownerAddress,
		OrderId:      orderId,
	}
}

func (msg *MsgCancelSpotOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid ownerAddress address (%s)", err)
	}

	// Validate order ID
	if msg.OrderId == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "order price cannot be 0")
	}

	return nil
}

var _ sdk.Msg = &MsgCancelSpotOrders{}

func NewMsgCancelSpotOrders(creator string, id []uint64) *MsgCancelSpotOrders {
	return &MsgCancelSpotOrders{
		SpotOrderIds: id,
		Creator:      creator,
	}
}

func (msg *MsgCancelSpotOrders) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	// Validate SpotOrderIds
	if len(msg.SpotOrderIds) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "spot order IDs cannot be empty")
	}
	for _, id := range msg.SpotOrderIds {
		if id == 0 {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "spot order ID cannot be zero")
		}
	}
	return nil
}
