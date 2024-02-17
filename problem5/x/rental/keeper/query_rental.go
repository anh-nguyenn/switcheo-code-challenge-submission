package keeper

import (
	"context"

	"rental/x/rental/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RentalAll(ctx context.Context, req *types.QueryAllRentalRequest) (*types.QueryAllRentalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var rentals []types.Rental

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	rentalStore := prefix.NewStore(store, types.KeyPrefix(types.RentalKey))

	pageRes, err := query.Paginate(rentalStore, req.Pagination, func(key []byte, value []byte) error {
		var rental types.Rental
		if err := k.cdc.Unmarshal(value, &rental); err != nil {
			return err
		}

		rentals = append(rentals, rental)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRentalResponse{Rental: rentals, Pagination: pageRes}, nil
}

func (k Keeper) Rental(ctx context.Context, req *types.QueryGetRentalRequest) (*types.QueryGetRentalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	rental, found := k.GetRental(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRentalResponse{Rental: rental}, nil
}
