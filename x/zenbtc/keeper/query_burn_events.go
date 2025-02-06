package keeper

import (
	"context"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryBurnEvents(ctx context.Context, req *types.QueryBurnEventsRequest) (*types.QueryBurnEventsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	burnEvents, err := k.BurnEvents.Get(ctx)
	if err != nil {
		return nil, err
	}

	matchingBurnEvents := make([]*types.BurnEvent, 0)

	for _, burnEvent := range burnEvents.Events[req.StartIndex:] {

		if (req.TxID == "" || burnEvent.TxID == req.TxID) &&
			(req.LogIndex == 0 || burnEvent.LogIndex == req.LogIndex) &&
			(req.ChainID == "" || burnEvent.ChainID == req.ChainID) {

			matchingBurnEvents = append(matchingBurnEvents, burnEvent)
		}
	}

	return &types.QueryBurnEventsResponse{BurnEvents: matchingBurnEvents}, nil
}
