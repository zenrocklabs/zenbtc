package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	"github.com/zenrocklabs/zenbtc/x/zenbtc/types"
)

func (k msgServer) UpdateParams(ctx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.GetParamsAuthority(ctx) != req.Authority {
		return nil, errorsmod.Wrapf(types.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetParamsAuthority(ctx), req.Authority)
	}

	if err := k.Params.Set(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
