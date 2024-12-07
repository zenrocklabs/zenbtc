package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func (k Keeper) Redemptions(goCtx context.Context, req *types.QueryRedemptionsRequest) (*types.QueryRedemptionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	redemptions := make([]types.Redemption, 0)
	var queryRange collections.Range[uint64]

	if err := k.validationKeeper.ZenBTCRedemptions.Walk(ctx, queryRange.StartInclusive(req.StartIndex), func(key uint64, redemption types.Redemption) (bool, error) {
		switch req.CompletionFilter {
		case types.CompletionFilter_COMPLETED:
			if redemption.Completed {
				redemptions = append(redemptions, redemption)
			}
		case types.CompletionFilter_PENDING:
			if !redemption.Completed {
				redemptions = append(redemptions, redemption)
			}
		default: // don't filter
			redemptions = append(redemptions, redemption)
		}
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &types.QueryRedemptionsResponse{Redemptions: redemptions}, nil
}
