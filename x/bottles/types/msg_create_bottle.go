package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBottle{}

func NewMsgCreateBottle(creator string, brand string, volume string) *MsgCreateBottle {
	return &MsgCreateBottle{
		Creator: creator,
		Brand:   brand,
		Volume:  volume,
	}
}

func (msg *MsgCreateBottle) Route() string {
	return RouterKey
}

func (msg *MsgCreateBottle) Type() string {
	return "CreateBottle"
}

func (msg *MsgCreateBottle) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBottle) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBottle) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Brand == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "brand cannot be empty")
	}
	return nil
}
