package keeper

import (
	"context"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bottles/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Bottle queries a specific bottle by its ID.
func (k Keeper) Bottle(c context.Context, req *types.QueryGetBottleRequest) (*types.QueryGetBottleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBottle(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "bottle not found")
	}

	return &types.QueryGetBottleResponse{Bottle: val}, nil
}

// BottleAll queries all bottles using your pagination function.
func (k Keeper) BottleAll(c context.Context, req *types.QueryAllBottleRequest) (*types.QueryAllBottleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	bottles, pageRes, err := k.GetBottlesWithPagination(ctx, req.Pagination)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBottleResponse{Bottle: bottles, Pagination: pageRes}, nil
}
