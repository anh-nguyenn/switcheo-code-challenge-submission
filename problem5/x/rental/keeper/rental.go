package keeper

import (
	"context"
	"encoding/binary"

	"rental/x/rental/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetRentalCount get the total number of rental
func (k Keeper) GetRentalCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.RentalCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRentalCount set the total number of rental
func (k Keeper) SetRentalCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.RentalCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRental appends a rental in the store with a new id and update the count
func (k Keeper) AppendRental(
	ctx context.Context,
	rental types.Rental,
) uint64 {
	// Create the rental
	count := k.GetRentalCount(ctx)

	// Set the ID of the appended value
	rental.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RentalKey))
	appendedValue := k.cdc.MustMarshal(&rental)
	store.Set(GetRentalIDBytes(rental.Id), appendedValue)

	// Update rental count
	k.SetRentalCount(ctx, count+1)

	return count
}

// SetRental set a specific rental in the store
func (k Keeper) SetRental(ctx context.Context, rental types.Rental) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RentalKey))
	b := k.cdc.MustMarshal(&rental)
	store.Set(GetRentalIDBytes(rental.Id), b)
}

// GetRental returns a rental from its id
func (k Keeper) GetRental(ctx context.Context, id uint64) (val types.Rental, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RentalKey))
	b := store.Get(GetRentalIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRental removes a rental from the store
func (k Keeper) RemoveRental(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RentalKey))
	store.Delete(GetRentalIDBytes(id))
}

// GetAllRental returns all rental
func (k Keeper) GetAllRental(ctx context.Context) (list []types.Rental) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RentalKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Rental
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRentalIDBytes returns the byte representation of the ID
func GetRentalIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.RentalKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
