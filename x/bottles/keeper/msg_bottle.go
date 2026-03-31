package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bottles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBottle(goCtx context.Context, msg *types.MsgCreateBottle) (*types.MsgCreateBottleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	count := k.GetBottleCount(ctx)

	var bottle = types.Bottle{
		Id:     count,
		Brand:  msg.Brand,
		Volume: msg.Volume,
		Owner:  msg.Creator,
	}

	k.SetBottle(ctx, bottle)

	k.SetBottleCount(ctx, count+1)

	return &types.MsgCreateBottleResponse{
		Id: count,
	}, nil
}
