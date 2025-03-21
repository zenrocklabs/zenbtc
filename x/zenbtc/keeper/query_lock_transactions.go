package keeper

import (
	"context"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func (k Keeper) LockTransactions(goCtx context.Context, req *types.QueryLockTransactionsRequest) (*types.QueryLockTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	lockTransactions := []*types.LockTransaction{}
	if err := k.LockTransactionStore.Walk(ctx, nil, func(key collections.Pair[string, uint64], value types.LockTransaction) (bool, error) {
		lockTransactions = append(lockTransactions, &value)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &types.QueryLockTransactionsResponse{LockTransactions: lockTransactions}, nil
}
