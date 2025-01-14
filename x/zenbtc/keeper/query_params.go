package keeper

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func (k Keeper) GetParams(ctx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	k.Logger().Warn("zenbtc", "params", fmt.Sprintf("%+v", k.Params))
	k.Logger().Warn("zenbtc", "ctx", fmt.Sprintf("%+v", ctx))
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: params}, nil
}
