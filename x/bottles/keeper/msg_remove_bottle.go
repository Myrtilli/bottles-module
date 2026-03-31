package keeper

import (
	"context"
	"fmt"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bottles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RemoveBottle(goCtx context.Context, msg *types.MsgRemoveBottle) (*types.MsgRemoveBottleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetBottle(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d not found", msg.Id))
	}

	if msg.Creator != val.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.Keeper.RemoveBottle(ctx, msg.Id)

	return &types.MsgRemoveBottleResponse{}, nil
}
