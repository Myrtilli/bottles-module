package keeper

import (
	"encoding/binary"

	"github.com/Bridgeless-Project/bridgeless-core/v12/x/bottles/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SetBottle saves bottle in store by its ID
func (k Keeper) SetBottle(ctx sdk.Context, v types.Bottle) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BottleKeyPrefix))
	b := k.cdc.MustMarshal(&v)
	store.Set(GetBottleIDBytes(v.Id), b)
}

// GetBottle returns bottle by its ID
func (k Keeper) GetBottle(ctx sdk.Context, id uint64) (val types.Bottle, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BottleKeyPrefix))
	b := store.Get(GetBottleIDBytes(id))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBottle remove bottle from the store
func (k Keeper) RemoveBottle(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BottleKeyPrefix))
	store.Delete(GetBottleIDBytes(id))
}

// GetBottlesWithPagination returns all Bottles with pagination
func (k Keeper) GetBottlesWithPagination(ctx sdk.Context, pagination *query.PageRequest) ([]types.Bottle, *query.PageResponse, error) {
	var bottles []types.Bottle
	store := ctx.KVStore(k.storeKey)
	bottleStore := prefix.NewStore(store, types.KeyPrefix(types.BottleKeyPrefix))

	pageRes, err := query.Paginate(bottleStore, pagination, func(key []byte, value []byte) error {
		var bottle types.Bottle
		k.cdc.MustUnmarshal(value, &bottle)

		bottles = append(bottles, bottle)
		return nil
	})

	if err != nil {
		return nil, nil, status.Error(codes.Internal, err.Error())
	}

	return bottles, pageRes, nil
}

func GetBottleIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBottleCount retrns all bottles count
func (k Keeper) GetBottleCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BottleCountKey))
	byteKey := []byte(types.BottleCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// SetBottleCount tracks count
func (k Keeper) SetBottleCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BottleCountKey))
	byteKey := []byte(types.BottleCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
