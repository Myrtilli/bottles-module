package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveBottle{}

func NewMsgRemoveBottle(creator string, id uint64) *MsgRemoveBottle {
	return &MsgRemoveBottle{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRemoveBottle) Route() string {
	return RouterKey
}

func (msg *MsgRemoveBottle) Type() string {
	return "RemoveBottle"
}

func (msg *MsgRemoveBottle) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveBottle) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveBottle) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
